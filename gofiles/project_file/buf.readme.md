# Preparando ambiente pra buf

## 1. Instalar set de ferramentas pra buf-cli desde golang

```bash
go install github.com/bufbuild/buf/cmd/buf@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
```
## 2. Configuración buf.yaml

Para más detalles, visita: https://buf.build/docs/configuration/v2/buf-yaml

**Ubicación actual:** `./buf.yaml`

```yaml
version: v2
modules:
  - path: grpcservice/proto
deps: 
  - buf.build/bufbuild/protovalidate
  - buf.build/googleapis/googleapis
  - buf.build/grpc-ecosystem/grpc-gateway
lint:
  use:
    - STANDARD
breaking:
  use:
    - FILE
```

### Parámetros principales:

- **modules**: Directorio donde se encuentran los archivos proto (`grpcservice/proto`)
- **deps**: Dependencias de protobuf que se utilizan
- **lint**: Linter estándar para validar los archivos proto
- **breaking**: Verificación de cambios incompatibles en la API
## 3. Configuración buf.gen.yaml

Para más detalles, visita: https://buf.build/plugins/protobuf

**Ubicación actual:** `./buf.gen.yaml`

```yaml
version: v2
plugins:
  # Plugin: Protocol Buffers para Go
  # Genera definiciones de mensajes proto para Go
  - remote: buf.build/protocolbuffers/go
    out: grpcservice/gen
    opt:
      - paths=source_relative
      - default_api_level=API_OPAQUE
      - Mprotoc-gen-openapiv2/options/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options

  # Plugin: Connect-RPC para Go
  # Genera definiciones de servicios (client/server) para Go
  - remote: buf.build/connectrpc/go
    out: grpcservice/gen
    opt:
      - paths=source_relative

  # Plugin: Protocol Buffers ES para Frontend
  # Genera mensajes (TypeScript/JavaScript) para el frontend
  - remote: buf.build/bufbuild/es
    out: grpcservice/gen/frontend
    opt: 
      - target=ts

  # Plugin: OpenAPI v2
  # Genera documentación OpenAPI
  - remote: buf.build/grpc-ecosystem/openapiv2:v2.27.3
    out: grpcservice/gen/openapi
    
managed:
  enabled: true
  override:
    - file_option: go_package_prefix
      value: ./ecs_govel/grpcservice/gen
  disable:
    - file_option: go_package 
      module: buf.build/bufbuild/protovalidate
    - module: buf.build/googleapis/googleapis
```

### Parámetros principales:

- **plugins**: Lista de generadores a utilizar
  - **protocolbuffers/go**: Genera código Go para mensajes
  - **connectrpc/go**: Genera código Go para servicios (cliente/servidor)
  - **bufbuild/es**: Genera código TypeScript para el frontend
  - **openapiv2**: Genera documentación OpenAPI

- **managed**: Configuración de opciones de archivo administradas
  - **go_package_prefix**: Prefijo del paquete Go para los archivos generados
  - **disable**: Deshabilita opciones para módulos específicos

## 4. Comandos para generar código

```bash
# Actualizar dependencias
$ buf dep update

# Validar archivos proto
$ buf lint

# Generar código
$ buf generate
```

## Estructura de directorios generados

```
grpcservice/gen/
├── <packages Go>/
│   ├── *.pb.go           (Mensajes compilados)
│   └── *.connect.go      (Servicios Connect-RPC)
├── frontend/
│   ├── *.ts              (Tipos TypeScript)
│   └── *.client.ts       (Cliente Connect-RPC TS)
└── openapi/
    └── *.openapi.yaml    (Documentación OpenAPI)
```