#!/bin/bash

path="./"

# Cargar las variables desde app.env (. -> carga variables ambiente)
. $path"app.env"

main_file="$APP_NAME"

echo "go build main.go"
go build -o $main_file
sleep 1

permision=$(stat -c "%a" "$path")
if [ $permision -ne 777 ]; then
    sudo chmod -R 777 $path
fi

if [ ! -d $path"Binary" ]; then
    mkdir -p $path"Binary"
    chmod -R 777 $path"Binary"
fi

permision=$(stat -c "%a" $path"Binary")
if [ $permision -ne 777 ]; then
    sudo chmod -R 777 $path"Binary"
fi

# Convertir la cadena en un arreglo
echo "$PORTS" | awk 'BEGIN {FS=","} {for (i=1; i<=NF; i++) print $i}' | while read -r port; do
    echo "$port"
    echo "killing $main_file$port"
    sudo pkill -f "$main_file$port"
    echo "copying $main_file to ${path}Binary/$main_file$port"
    #cp $main_file $path"Binary/$main_file$port"
    rsync -av --inplace $main_file $path"Binary/$main_file$port"
done

rm $main_file
