# El pacs apidocs solo necesita la firma de la función y los comentarios de sawggo para generar la api docs de swagger
<details>
  <summary>Ejemplo de creación de api docs tipo swagger</summary>

```
    //	@Schemes http https
    // @Tags Routes of test
    // @Accept json
    // @Produce json
    // @Description ruta de test
    // @Success 200 string OK
    // @Failure 400 {object} object
    // @Failure 404 {object} object
    // @Router /test [get]
    func RutaTestFunc(c *gin.Context) {} // esta es la firma del handler de ejemplo, solo se usa para generar la documentación
```
</details>