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
