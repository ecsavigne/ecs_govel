---
applyTo: "**/*.{vue,ts,js}"
---

# Instrucciones de Frontend

## Conciencia de Contratos
- Las APIs del backend son contratos estrictos
- Nunca asumas campos que no estén definidos en protobuf
- Trata los campos faltantes u opcionales de forma defensiva

## Arquitectura
- Los componentes de UI deben centrarse en la presentación
- La lógica de negocio pertenece a servicios o composables
- Evita acoplar la UI a detalles internos del backend

## Manejo de Datos
- Usa clientes de API fuertemente tipados
- Valida y normaliza los datos de la API en los límites del sistema
- Maneja explícitamente los estados de carga, vacío y error

## Errores
- Muestra mensajes claros y amigables al usuario
- Nunca ignores errores de API o de parsing
