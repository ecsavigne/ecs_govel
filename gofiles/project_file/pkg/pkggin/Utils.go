package pkggin

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"strings"

	c_ "ecs_govel/configs"
	"ecs_govel/pkg/pkglog"
	"ecs_govel/pkg/pkgutil"

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

func secureServer(route *gin.Engine) {
	expectHost := HostAllow()

	route.Use(func(c *gin.Context) {
		host := pkgutil.ParseHost(c.Request.Host)
		if _, ok := slices.BinarySearch(expectHost, host); !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("invalid host expected: %s, received: %s", expectHost, c.Request.Host)})
			return
		}

		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		c.Next()
	})
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

	if c_.StateInitMetric {
		//TODO: Configurar el motor de Gin para metricas no esta implementada la logica aun
		// fmt.Println("Configurar el motor de Gin para metricas y documentacion")
		pkglog.Log.Sub("Configs").Infof("Begin the engine of Gin for metricas\n")
		metricEngine = gin.Default()
		secureServer(metricEngine)
	}

	if c_.StateInitDocApi {
		//TODO: Configurar el motor de Gin para metricas no esta implementada la logica aun
		// fmt.Println("Configurar el motor de Gin para metricas y documentacion")
		pkglog.Log.Sub("Configs").Infof("Begin the engine of Gin for Documents\n")
		docApiEngine = gin.Default()
		secureServer(docApiEngine)
	}

	engine = gin.New()
	secureServer(engine)

	gin.ForceConsoleColor()
	engine.Use(gin.Recovery(), gin.Logger())
	engine.MaxMultipartMemory = 100 << 20
}
