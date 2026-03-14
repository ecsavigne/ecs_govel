package apidoc

import (
	"github.com/gin-gonic/gin"
)

//	@Schemes http https
//
// @Tags Routes of test
// @Accept json
// @Produce json
// @Description ruta de test
// @Success 200 string OK
// @Failure 400 {object} object
// @Failure 404 {object} object
// @Router /test [get]
func RutaTestFunc(c *gin.Context) {}
