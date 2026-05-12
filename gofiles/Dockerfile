# instalar sistema de compilação
FROM amd64/ubuntu:noble
# FROM golang:latest

#Instalacion de
RUN apt update && apt install -y curl wget tar grep rsync

# install supervisor
RUN apt-cache show supervisor && apt update && apt install -y supervisor
RUN supervisord -v

RUN  /usr/bin/echo_supervisord_conf > /etc/supervisor/supervisord.conf && \
    echo "\n[include]\nfiles = /etc/supervisor/conf.d/*.conf" >> /etc/supervisor/supervisord.conf
RUN ln -s /etc/supervisor/supervisord.conf /etc/supervisord.conf

# # Preparar diretório de trabajo
WORKDIR /app

# # instalar golang
#     # descargar la ultima version
RUN VERSION=$(curl -s https://go.dev/VERSION?m=text | grep -oE '^go[0-9]+\.[0-9]+\.[0-9]+') && \
    echo "Go version: $VERSION" && \
    wget https://go.dev/dl/$VERSION.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf $VERSION.linux-amd64.tar.gz && \
       rm $VERSION.linux-amd64.tar.gz
  
#Configurar variables de Entorno
# ENV PATH ${PATH}:/usr/local/go/bin
ENV PATH="/usr/local/go/bin:${PATH}"

# files of project( project github)
COPY project_file/ .

# Executar compilacion
# RUN ./build.sh