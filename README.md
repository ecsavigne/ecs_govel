# [Inspirado en laravel](#lintest)
<a name="lintest"></a>
1- Para probara si el server esta corriendo segun la configuracion de ip y proxy creada en
# config/app.ini
```
Aqui se encuentra la configuracion correspondiente al servidor Web y a la App
```
# db.ini 
```
Aqui ó, correspondiente a base datos, driver path del archivo env con variable
```
## var para base de datos ecs_govel:
 ej: Config archivo .env para base de datos driver=mysql
```
 DB_HOST=localhost
 DB_PORT=${MYSQL_FORWARD_PORT}
 DB_USER=${MYSQL_USER}
 DB_NAME=${MYSQL_DATABASE}
 DB_PASSWD=${MYSQL_PASSWORD}
```

# [Api Doc](#lintest1)
<a name="lintest1"></a>
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
