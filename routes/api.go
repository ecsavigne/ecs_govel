package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Schemes http https
//
// @Tags Routes of test
// @Accept json
// @Produce json
// @Description ruta de test
// @Success 200 {object} object
// @Failure 400 {object} object
// @Router /test [post]
func RutaTestFunc(c *gin.Context) {
	fmt.Print("RutaTestFunc")
	c.String(http.StatusOK, "RutaTestFunc")
}

func RutaTestFunc1(c *gin.Context) {
	fmt.Print("RutaTestFunc")
	// Response + :Param, ?Query, Data POST, Form POST, Multipart POST
	c.JSON(http.StatusOK, gin.H{
		"Param 1":      c.Param("name"),                     // Con parametro router.GET("/user/:id", function(){})
		"Query2(name)": c.Query("name"),                     // Con Query GET /path?id=1234&name=Manu&value
		"ValorTest1":   c.Request.FormValue("name"),         //Deprecated Desde Formulario(Form-Data)
		"T":            c.PostForm("t"),                     //Ok urlencoded form or multipart form when it exists
		"T1":           c.DefaultPostForm("Name", "asassd"), //urlencoded form or multipart form when it exists si no coloca valor por defecto
	})

	// -------Guardando binario file
	// -------Obtenerlo desde la peticion
	// _, fileHeader, err := c.Request.FormFile("file")
	// if err != nil {
	// 	c.String(http.StatusBadRequest, err.Error())
	// 	return
	// }
	// --- Salva el archivo en el disco
	// if err := c.SaveUploadedFile(fileHeader, dst); err != nil {
	// 	c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
	// 	return
	// }
}

func loadApi(g *gin.Engine) {
	// Sin parametro
	g.POST("/Test", RutaTestFunc)
}
