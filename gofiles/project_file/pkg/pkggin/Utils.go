package pkggin

import (
	"fmt"
	"io"
	"os"
	"strings"

	c_ "ecs_govel/configs"
	"ecs_govel/pkg/pkglog"

	"github.com/gin-gonic/gin"
)

func RouterList(g *gin.Engine) {
	fmt.Printf("\n\n%d Rutas listas en %s:%s \n\n", len(g.Routes()), c_.HTTP_SERVER_HOST, c_.HTTP_SERVER_PORT)
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("| Method\t | Path\t\t\t\t")
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	i := 0
	for _, route := range g.Routes() {
		fmt.Println("Ruta #", i+1)
		i++
		switch route.Method {
		case "GET":
			fmt.Printf("| \033[33m%s\033[0m\t\t | \033[32m%s\033[0m\t\t \n", route.Method, route.Path)
		case "POST":
			fmt.Printf("| \033[34m%s\033[0m\t\t | \033[32m%s\033[0m\t\t \n", route.Method, route.Path)
		case "PUT":
			fmt.Printf("| \033[35m%s\033[0m\t\t | \033[32m%s\033[0m\t\t \n", route.Method, route.Path)
		case "PATCH":
			fmt.Printf("| \033[36m%s\033[0m\t\t | \033[32m%s\033[0m\t\t \n", route.Method, route.Path)
		case "DELETE":
			fmt.Printf("| \033[31m%s\033[0m\t | \033[32m%s\033[0m\t\t \n", route.Method, route.Path)
		default:
			fmt.Printf("| \033[32m%s\033[0m\t\t | \033[32m%s\033[0m\t\t \n", route.Method, route.Path)
		}
		fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	}
}

func GetEngine() *gin.Engine {
	return engine
}

func GetMetricEngine() *gin.Engine {
	return metricEngine
}

func GetDocApiEngine() *gin.Engine {
	return docApiEngine
}

// Configurar el motor de Gin
func prepare_engine() {
	if strings.ToLower(c_.APP_MODE) == "develop" {
		fmt.Println("DEVELOP MODE")
		gin.SetMode(gin.DebugMode)
	} else {
		fmt.Println("RELEASE MODE")
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.MultiWriter(c_.GetFileLogger(), os.Stdout)
	}

	if c_.IsX1() {
		//TODO: Configurar el motor de Gin para metricas no esta implementada la logica aun
		// fmt.Println("Configurar el motor de Gin para metricas y documentacion")
		pkglog.Log.Sub("Configs").Infof("Configurar el motor de Gin para metricas y documentacion\n")
		docApiEngine = gin.Default()

		metricEngine = gin.Default()
	}

	engine = gin.Default()
	engine.MaxMultipartMemory = 100 << 20
}
