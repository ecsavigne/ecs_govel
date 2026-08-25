# instalar sistema de compilação
FROM amd64/ubuntu:noble
# FROM golang:latest

# #Instalacion de
# RUN apt update && apt install -y curl wget tar grep rsync gcc libc6-dev

# # install supervisor
# RUN apt-cache show supervisor && apt update && apt install -y supervisor
# RUN supervisord -v

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        curl \
        wget \
        tar \
        grep \
        rsync \
        gcc \
        libc6-dev \
        ca-certificates \
        supervisor && \
        update-ca-certificates

RUN  /usr/bin/echo_supervisord_conf > /etc/supervisor/supervisord.conf && \
    echo "\n[include]\nfiles = /etc/supervisor/conf.d/*.conf" >> /etc/supervisor/supervisord.conf
RUN ln -s /etc/supervisor/supervisord.conf /etc/supervisord.conf

# # instalar golang
#     # descargar la ultima version
RUN curl -fsSL 'https://go.dev/VERSION?m=text' \
    | grep -m1 -oE '^go[0-9]+\.[0-9]+\.[0-9]+' \
    > /tmp/go_version \
    && export GO_VERSION="$(cat /tmp/go_version)" && \
    echo "Go version: ${GO_VERSION}" && \
    wget "https://go.dev/dl/${GO_VERSION}.linux-amd64.tar.gz" -O /tmp/go.tar.gz
# clean space
RUN rm -rf /usr/local/go && \
    tar -C /usr/local -xzf /tmp/go.tar.gz && \
    rm -f /tmp/go.tar.gz /tmp/go_version && \
    rm -rf /var/lib/apt/lists/*
  
#Configurar variables de Entorno
ENV PATH="/usr/local/go/bin:${PATH}"
ENV CGO_ENABLED=1

# # Preparar diretório de trabajo
WORKDIR /app

# files of project( project github)
COPY project_file/ .