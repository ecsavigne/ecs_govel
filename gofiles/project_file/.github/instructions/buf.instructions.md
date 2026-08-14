---
applyTo: "**/{buf.yaml,buf.gen.yaml,buf.work.yaml}"
---

# Instrucciones de Buf
- Para preparar el ambiente si no esta instalado vas usar las instruciones del archivo `buf.prepeare.md`

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

## Checklist

- Cuando actualice dependencia de buf debe ser con `buf dep update`
-  mostrar cuales son los mensajes a generar y las estructuras de golang antes de crear, luego usuario **decide**
-  buf.gen.yaml no puede tener `include_imports: true`
-  los archivos `buf.yaml y buf.gen.yaml` deben seguir estrictamente las instruciones dada aqui.
-  Si se importa algun mensaje dentro de otro asegurase de que sea creadas correctamente las opciones del plugin de generacion para hacer la importacion correctamente.
-  la session de 
  ```manager  override:
    - file_option: go_package_prefix
    value: ...
   ``` 
   incluya el path obsoluto del modulo
-  Asegurate de que cada checklist se cumpla antes de generar y crear codigos.