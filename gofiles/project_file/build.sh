#!/bin/bash

path="./"

# Cargar las variables desde app.env (. -> carga variables ambiente)
. $path"app.env"

main_file=${SERVICE_NAME}
folder_exec=${FOLDER_EXEC}

module_name=ecs_govel

go mod init ${module_name} || echo "Módulo ya inicializado"
if [ ! -d "vendor" ]; then
    go mod tidy
    go mod vendor
fi 

if [ ${UPDATE_DEPENDENCIES} == true ]; then
    go get -u
    go mod vendor
    echo "update dependencies..."
fi

echo "go build file: '${main_file}'"
go build -o ${main_file}
sleep 1

permision=$(stat -c "%a" "${path}")
if [ $permision -ne 777 ]; then
    # for docker
    chmod -R 777 ${path}
fi

if [ ! -d $path"${folder_exec}" ]; then
    mkdir -p ${path}${folder_exec}
    chmod -R 777 ${path}${folder_exec}
fi

permision=$(stat -c "%a" $path"${folder_exec}")
if [ $permision -ne 777 ]; then
    # for docker
    chmod -R 777 ${path}${folder_exec}
fi

FS=',' read -ra ADDR <<< "${PORTS}"
for port in "${ADDR[@]}"; do
    # Limpiar espacios en blanco si los hubiera
    port=$(echo "$port" | xargs)
    [ -z "$port" ] && continue

    echo "${port}"
    echo "killing ${main_file}${port}"
    pkill -f "${main_file}${port}" || true

    echo "copying ${main_file} to ${path}${folder_exec}/${main_file}${port}"
    rsync -av --inplace "${main_file}" "${path}${folder_exec}/${main_file}${port}"
done

# Eliminar el ejecutable original de forma segura
if [ -n "${main_file}" ] && [ -f "${main_file}" ]; then
    rm -f "${main_file}"
fi

supervisord -c /etc/supervisor/supervisord.conf
./cmd/${main_file}${port}