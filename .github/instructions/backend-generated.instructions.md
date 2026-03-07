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