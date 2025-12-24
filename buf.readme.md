# Preparando ambiente pra buf
### 1. Instalar set de ferramentas pra buf-cli desde golang
        $ mkdir connect-go-example
        $ cd connect-go-example
        $ go mod init example
        $ go install github.com/bufbuild/buf/cmd/buf@latest
        $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
        $ go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
### 2. Criar buf.yaml
        # For details on buf.yaml configuration, visit https://buf.build/docs/configuration/v2/buf-yaml
        version: v2
        modules: # agrega dir de modulos por defecto donde se van a buscar los protos
            - path: crud_mongo_db/proto
        # Dependenças dos protos que vão se utilizar. ver: https://buf.build/explore
        deps: 
        - buf.build/bufbuild/protovalidate
        - buf.build/googleapis/googleapis
        - buf.build/grpc-ecosystem/grpc-gateway
        lint:
        use:
            - STANDARD
        #   except: ignorar reglas de lint
            # - PACKAGE_DIRECTORY_MATCH    # Ignora que la carpeta no coincida
            # - PACKAGE_VERSION_SUFFIX      # Ignora que no tenga .v1
            # - RPC_REQUEST_STANDARD_NAME   # Ignora el nombre de los mensajes Request
            # - RPC_RESPONSE_STANDARD_NAME  # Ignora el nombre de los mensajes Response  
        breaking:
        use:
            - FILE
### 3. Criar <i>buf.ge.yaml</i>
        version: v2
        # Ver mais info about plugins: https://buf.build/plugins/protobuf
        # github use and docs: https://github.com/bufbuild
        plugins:
        # Gera definição pra os mensagens proto para go
        # github use and docs: https://github.com/protocolbuffers
        # url plugin: https://buf.build/protocolbuffers/go
        - remote: buf.build/protocolbuffers/go
            out: crud_mongo_db/gen
            opt:
            - paths=source_relative
            # values: API_OPAQUE, API_HYBRID, API_OPEN  
            - default_api_level=API_OPAQUE
            - Mprotoc-gen-openapiv2/options/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options

        # Gera definição pra os servicos (client/server connect-rpc para go)
        # github use and docs: https://github.com/connectrpc/connect-go
        # url plugin: https://buf.build/connectrpc/go
        - remote: buf.build/connectrpc/go
            out: crud_mongo_db/gen
            opt:
            - paths=source_relative
            - simple

        # Genera los mensajes (clases de TS/JS y client/server connect-rpc) para el frontend
        # github use and docs: https://github.com/bufbuild/protobuf-es
        # url plugin: https://buf.build/bufbuild/es
        - remote: buf.build/bufbuild/es
            out: crud_mongo_db/gen/frontend
            # include_imports: true
            opt: 
            - target=ts # Puedes usar 'ts' o 'js'

        # Documentacion openapi  url: https://github.com/google/gnostic/tree/main/cmd/protoc-gen-openapi
        - remote: buf.build/grpc-ecosystem/openapiv2:v2.27.3
            out:  crud_mongo_db/gen/openapi
            
        managed:
        enabled: true
        override:
            - file_option: go_package_prefix
            value: ./services_ej_1/crud_mongo_db/gen
        disable:
            - file_option: go_package 
            module: buf.build/bufbuild/protovalidate
            - module: buf.build/googleapis/googleapis

### 4. Gerar código
        $ buf dep update
        $ buf lint
        $ buf generate