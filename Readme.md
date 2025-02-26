# Api Doc
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
