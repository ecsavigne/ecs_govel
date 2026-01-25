---
applyTo: "**/internal/**/*.go"
---

# Instrucciones para Código Backend Escrito a Mano

## Arquitectura
- El código generado define interfaces y contratos
- El código escrito a mano implementa únicamente la lógica de negocio
- Nunca modifiques las interfaces generadas

## Buenas Prácticas en Go
- Sigue patrones idiomáticos de Go
- Retorna errores de forma explícita
- Mantén los handlers livianos; mueve la lógica a servicios

## gRPC
- Valida las entradas de forma explícita
- Mapea los errores de dominio a códigos de estado gRPC
- No expongas errores internos a los clientes

## Seguridad
- Asume que los handlers pueden ejecutarse de forma concurrente
- Evita estado mutable compartido