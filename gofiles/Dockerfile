# instalar sistema de compilação
FROM amd64/ubuntu:noble
# FROM golang:latest

#Instalacion de
RUN apt update && apt install -y curl wget tar

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
COPY project_file/* .

# Executar
CMD ["sh","-c", "ls -la /app & tail -f /dev/null"]