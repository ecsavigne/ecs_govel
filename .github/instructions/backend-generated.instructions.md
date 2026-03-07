---
applyTo: "**/gen/**"
---

# Instrucciones para Código Backend Generado

## Política de Código Generado
- Este código se genera automáticamente a partir de protobuf usando buf
- Nunca edites manualmente los archivos generados
- Cualquier cambio debe realizarse en los archivos `.proto`

## Restricciones
- No refactorices
- No reformatees
- No agregues lógica, comentarios ni funciones auxiliares

## Regeneración
- Asume que este directorio puede eliminarse y regenerarse en cualquier momento

# Checklist de Revisión
- [ ] El código implementa solo la lógica de negocio, sin modificar las interfaces generadas
- [ ] el cliente gRPC debe seguir la estructura `grpcservice/app/client/client.go.ej`
- [ ] El servidor gRPC debe seguir la estructura `grpcservice/app/server/server.go.ej`
- [ ] A la hora de crear `g.Use(otelgin.Middleware("nombre-del-servicio"))` en el servidor gRPC ej:

```proto
service ShortURLService {
  // RegisterUser crea una nueva cuenta de usuario en el sistema.
  // Solo usuarios únicos por email son permitidos.
  // La contraseña se almacena hasheada con bcrypt (cost 12).
  // Casos de uso: RF-01
  // Errores: ALREADY_EXISTS (email duplicado), INVALID_ARGUMENT (validaciones)
  rpc RegisterUser(RegisterUserRequest) returns (RegisterUserResponse) {
    option (google.api.http) = {
      post: "/shorturl/v1/auth/register"
      body: "*"
    };
    option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
      summary: "Registrar usuario (RF-01)"
      description: "Crea una cuenta nueva. El email debe ser único. La contraseña se almacena hasheada, nunca en texto plano."
      tags: ["Autenticación"]
    };
  }
} 
```

```go file: grpcservice/app/server/server.go

func InitGrpcService() {
  g.Use(otelgin.Middleware("shorturl"))

/*
part code
*/

path, handler := conn.NewShortURLServiceHandler(
		new(ShortURLServiceImpl),
		// Validation via Protovalidate is almost always recommended
		interceptors,
	)

  /*
part code
*/

// routes for anotations proto
	routerGin.Any("/shorturl/*any", gin.WrapH(transcoder))
}
  ```

y en el cliente gRPC ej:

- [ ]  La llamada en el main debe seguir la estructura:
```go file: main.go

// Init GrpcService
func run() {
/*part code*/
	if strings.ToLower(config.TYPE_SERVICES) == "grpc" {
		grpcserverinit.InitGrpcService()
	} else {
		wait := make(chan os.Signal, 1)
		signal.Notify(wait, os.Interrupt, syscall.SIGTERM)
		<-wait
	}

  /*part code*/
}  
```
- [ ] El código generado no debe tener errores de compilación ni advertencias
- [ ] Solo debe existir una sola vez esto `routerGin.Any("/shorturl/*any", gin.WrapH(transcoder))` ej:
```go file: grpcservice/app/server/server.go


  routerGin.Any("/productsapi/*any", gin.WrapH(transcoder)) // Mal
  routerGin.Any("/shorturl/*any", gin.WrapH(transcoder))
```
- [ ] El nombre del paquete `grpcservice/app/server` debe ser `server` y el nombre del paquete `grpcservice/app/client` debe ser `client`
- [ ] la var ambiente `DOC_API_PATH` muestra la ruta correcta del archivo `name_service(no camel case).swagger.json` generado por buf ej:
```bash
  # ShortURLService -> shorturl_service.swagger.json
  DOC_API_PATH=./grpcservice/gen/openapi/services/v1/shorturl_service.swagger.json
```