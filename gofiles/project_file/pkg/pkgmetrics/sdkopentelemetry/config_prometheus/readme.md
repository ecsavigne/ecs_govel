# Configuração de prometheus
## Multi serviços. Opción 1: Si los servicios son distintos (Recomendado)
        scrape_configs:
        # Servicio de Productos
        - job_name: 'product-service'
            static_configs:
            - targets: ['host.docker.internal:9464']

        # Servicio de Usuarios (corriendo en otro puerto)
        - job_name: 'user-service'
            static_configs:
            - targets: ['host.docker.internal:9465']

        # Servicio de Inventario
        - job_name: 'inventory-service'
            static_configs:
            - targets: ['host.docker.internal:9466']
### O porto vai ser igual ao de server de metricas
## Opción 2: Si es el mismo servicio pero con varias réplicas
          - job_name: 'product-service-cluster'
      static_configs:
        - targets: 
            - 'host.docker.internal:9464'
            - 'host.docker.internal:9465'
            - 'host.docker.internal:9466'
# Grafico de grafana
1. id: 14765 -> Ideal pra mostrar metricas de serviços grpc