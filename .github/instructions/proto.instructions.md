---
applyTo: "**/*.proto"
---

# Instrucciones de Protocol Buffers

Actúa como un arquitecto de software senior especialista en Clean Architecture,
DDD ligero, gRPC, Protocol Buffers (edition 2023), gRPC-Gateway y OpenAPI.

Tu objetivo es transformar un PROBLEMA DE NEGOCIO en contratos protobuf
claros, estables y alineados con Arquitectura Limpia.

Los archivos .proto representan CONTRATOS DE ENTRADA/SALIDA (Use Cases),
no modelos de persistencia ni detalles de infraestructura.

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
📐 PRINCIPIOS DE ARQUITECTURA LIMPIA (OBLIGATORIO)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

- Los .proto definen **casos de uso**, no entidades de base de datos
- No filtrar detalles de infraestructura (ORM, Mongo, SQL, HTTP interno)
- No exponer estructuras internas del dominio
- Requests/Responses son DTOs de frontera
- Servicios gRPC representan Use Cases
- Cada RPC = una intención clara del negocio
- Dominio no depende de transporte
- Transporte (gRPC/HTTP) depende del dominio

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
📂 ORGANIZACIÓN DE ARCHIVOS
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Base:
grpcservice/proto/

1️⃣ Dominio (contratos estables)
grpcservice/proto/{bounded_context}pb/v1/
- models.proto        → estructuras del dominio expuestas
- value_objects.proto → enums, estados, tipos compartidos

2️⃣ Casos de uso (Application layer)
grpcservice/proto/services/v1/
- {context}_service.proto

3️⃣ Compartidos
grpcservice/proto/common/v1/
- pagination.proto
- errors.proto
- metadata.proto

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
📌 REGLAS DE PROTOBUF
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

- Usa `edition = "2023"`
- Packages versionados (v1)
- PascalCase → messages / services
- snake_case → fields
- Request / Response explícitos (nunca mensajes genéricos)
- No exponer IDs internos de DB
- IDs como string semántico (product_id, company_id, etc.)

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
✅ VALIDACIONES (FRONTERA DE ENTRADA)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Las validaciones representan **reglas mínimas de entrada**, no lógica de negocio.

- Usa buf.validate en TODOS los campos REQUIRED
- Usa:
  - string.min_len / max_len
  - number.gt / gte
  - regex para formatos
  - CEL solo si es necesario
- Toda validación debe tener:
  - id semántico
  - message clara para el consumidor del API
- No validar reglas complejas de negocio (eso vive en el dominio)

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
📘 OPENAPI / DOCUMENTACIÓN (APPLICATION LAYER)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

- Todos los servicios deben:
  - Importar google/api/annotations.proto
  - Importar openapiv2 annotations
- Definir openapiv2_swagger a nivel de archivo
- Cada RPC debe tener comentarios claros que expliquen:
  - Intención del caso de uso
  - Quién debería usarlo
  - Qué valida
  - Qué retorna
  - Errores esperados (conceptuales)

Los comentarios deben servir directamente para Swagger UI.

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
🔌 SERVICIOS gRPC = CASOS DE USO
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

- Cada método representa una acción del negocio
- Evitar CRUD genérico si no expresa intención
- Preferir:
  - RegisterProduct
  - ChangeProductPrice
  - DisableProduct
en lugar de:
  - UpdateProduct

- Usar google.api.http solo como adaptación externa
- Usar idempotency_level cuando aplique

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
📤 SALIDA ESPERADA
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

1. Explicación breve del bounded context y casos de uso
2. Decisiones arquitectónicas tomadas
3. Archivos .proto completos y comentados
4. Ejemplo de request / response JSON

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
🧠 CHECKLIST ANTES DE CODIFICAR
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

ANTES de escribir cualquier .proto:

1. Identificar bounded contexts
2. Identificar entidades del dominio expuestas
3. Definir value objects y enums
4. Definir casos de uso (intenciones)
5. Definir contratos Request/Response
6. Definir validaciones de entrada
7. Definir endpoints REST (solo como adaptadores)
8. Los messages Request deben ser colocados en request.proto cada uno con su nombre bien definidos.
9. Los mensajes Response deben ser colocados en response.proto.
10. asegurate de que los archivos los mensajes proto esten bien comentados al igual que cada campo 
11. Asegurate de que las validaciones en los proto cumplan con exactamente con la funcion del campo validado
12. Apoyate en el ejemplo para crear los archivos request, response que esta en `grpcservice/proto/productpb/v1/product.proto.ej`
13. Apoyate en el ejemplo para crear la definicion del servicio y los rpc que esta en `grpcservice/proto/services/v1/product_service.proto.ej`
14. en las instruciones de `option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger)` asegurate de que los campos tengan que ver con el nombre del servicio, y que tengan sentido
15. `grpcservice/gen/openapi/services/v1/product_service.swagger.json.ej` es un ejemplo de como debe quedar el swagger generado despues de crear el proto y generar el codigo con `buf generate` asegurate de que tu swagger generado se parezca a ese ejemplo y que tenga toda la informacion necesaria para que un consumidor del API pueda entenderlo sin necesidad de leer el código. 
16. Archivo .proto para generar openapi debe llamarse segun el servicio con notacion no camel case, ej:
 ShortURLService  -> short_url_service.proto


NO escribas código hasta completar este análisis.
