# [Inspirado en laravel](#intro)
<a name="intro"></a>
1- Para probara si el server esta corriendo segun la configuracion de ip y proxy creada en
# config/app.ini
```
Aqui se encuentra la configuracion correspondiente al servidor Web y a la App
```
# app.env
```
Aqui ó, correspondiente a base datos, driver path del archivo env con variable
```
## var para configuracion separada por sessiones:
 ej: Config archivo app.env para base de datos driver=mysql
```
# Database
FORWARD_DB_PORT=5432
PG_DB_PORT=5432
PG_DB_HOST=localhost
PG_DB_DATABASE=postgres
PG_DB_USERNAME=postgres
PG_DB_PASSWORD=password
PG_DNS_LOCAL=host=$PG_DB_HOST port=$PG_DB_PORT user=$PG_DB_USERNAME password=$PG_DB_PASSWORD sslmode=disable

#APP
MAX_CONNECTIONS=15
CANT_X=1
FILE_LOGGER=/var/log/whatsmeow/all_logs2.log
SESSIONS=/var/www/html/whatsmeow_storage/sessions
MESSAJE_FILES=/var/www/html/whatsmeow_storage/storage
AVATAR_FILES=/var/www/html/whatsmeow_storage/avatars
NAME_X1=main13371

#server http
HTTP_SERVER_HOST=localhost
HTTP_SERVER_HOST_METRICS=localhost
HTTP_SERVER_PORT_METRICS=1111
HTTP_SERVER_PORT_TEST=3333
HTTP_SERVER_PORT_DOC_API=2222
```

# [Api Doc](#apiDoc)
<a name="apiDoc"></a>
## Para Correr api doc hay que instalar:
### 1.
 ```
    $ go install github.com/swaggo/swag/cmd/swag@latest
    $ go get -u github.com/swaggo/gin-swagger
    $ go get -u github.com/swaggo/files
 ```
 ### 2. Colocar los comentarios en las rutas a documentar
 ### 3. Ejecutar comando para generar achivos de docs en folder docs:
 ```
    $ swag init
 ``` 
 ### Ejemplo de documentacion:
 ```
    /* Mas info: https://github.com/swaggo/gin-swagger,  geral: https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format*/
    // @BasePath /api/v1

    // PingExample godoc
    // @Summary ping example
    // @Schemes
    // @Description do ping
    // @Tags example
    // @Accept json
    // @Produce json
    // @Success 200 {string} Helloworld
    // @Router /example/helloworld [get]
    func Helloworld(g *gin.Context)  {
        g.JSON(http.StatusOK,"helloworld")
    }

 ```
# [Opcional] Para activar recarga live del código en desarrollo se debe seguir estos pasos:
## 1. Install air, se debe estar en la carpeta principal de ecs_govel
```
   go install github.com/air-verse/air@latest
```
## 2. Inicializar air
```
   air init
```
## 3. Ejecutar air
```
   air
```
### Saber más sobre [air](#air)
<a name="air" href="https://github.com/air-verse/air" target="_blank"></a>