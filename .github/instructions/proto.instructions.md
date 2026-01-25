---
applyTo: "**/*.proto"
---

# Instrucciones de Protocol Buffers

## Fuente de la Verdad
- Los archivos `.proto` definen la API canónica
- Backend y frontend deben seguir estrictamente las definiciones de protobuf

## Reglas de Diseño
- Los nombres de mensajes y servicios deben ser estables y explícitos
- Nunca reutilices ni modifiques los números de campo
- Prefiere agregar nuevos campos en lugar de modificar los existentes

## Compatibilidad
- Evita cambios incompatibles
- Marca los campos obsoletos como `deprecated` en lugar de eliminarlos
- Los nuevos campos deben ser opcionales siempre que sea posible

## Documentación
- Agrega comentarios cuando la intención no sea evidente
- Asume que las APIs son consumidas por sistemas externos
