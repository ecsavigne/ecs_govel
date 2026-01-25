---
applyTo: "**/{buf.yaml,buf.gen.yaml,buf.work.yaml}"
---

# Instrucciones de Buf

## Buf como Autoridad
- Asume que `buf` gestiona la generación, el linting y las verificaciones de cambios incompatibles
- No eludas ni omitas las reglas definidas por buf

## Flujo de Trabajo
- Los cambios en la API deben seguir este orden:
  1. Actualizar los archivos `.proto`
  2. Validar con `buf lint`
  3. Verificar compatibilidad con `buf breaking`
  4. Generar el código con `buf generate`

## Seguridad
- No debilites las reglas de lint ni de breaking
- Trata los fallos como problemas de diseño, no como obstáculos
