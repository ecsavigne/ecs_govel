package configs

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/ecsavigne/proxy-reverse/proxy"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	routerApi                 *gin.Engine
	routerMetric              *gin.Engine
	routerDocApi              *gin.Engine
	APP_MAX_CONNECTIONS       int
	APP_CANT_X                int
	APP_FILE_LOGGER           string
	APP_SESSIONS              string
	APP_MESSAJE_FILES         string
	APP_AVATAR_FILES          string
	APP_NAME_X1               string
	APP_PROCESS_EVENT_RECIVED bool
)

func init() {
	prepare_app()
}

func GetEngine() *gin.Engine {
	return routerApi
}

func GetMetricEngine() *gin.Engine {
	return routerMetric
}

func GetDocApiEngine() *gin.Engine {
	return routerDocApi
}

func prepare_proxies() {
	ReverseProxyApi = proxy.NewProxyReverse(proxy.ProxyReverse{
		Host: HTTP_SERVER_HOST,
		Port: HTTP_SERVER_PORT,
	})
}

// Configurar el motor de Gin
func prepare_engine() {
	gin.SetMode(gin.ReleaseMode)
	routerApi = gin.Default()
	if IsX1() {
		//TODO: Configurar el motor de Gin para metricas no esta implementada la logica aun
		routerMetric = gin.Default()
		routerDocApi = gin.Default()
	}
}

// Carga de varEnv
func prepare_env() {
	pathDir, _ := os.Getwd()
	viper.AddConfigPath(pathDir)
	viper.SetConfigType("env")
	viper.SetConfigName("app.env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("\033[31mError: No encontrado archivo app.env ni .cobraToml de tipo (toml) en\033[30m %s\n", pathDir)
		os.Exit(2)
	} else {
		fmt.Println("-----------------------------------------------------------")
		fmt.Println("\033[32mEncontrado archivo tipo\033[0m \033[34m(env)\033[0m")

		// variables .env DB
		DB_HOST = viper.GetString("PG_DB_HOST")
		DB_USER = viper.GetString("PG_DB_USERNAME")
		DB_NAME = viper.GetString("PG_DB_DATABASE")
		DB_PASSWORD = viper.GetString("PG_DB_PASSWORD")
		DB_PORT = viper.GetString("PG_DB_PORT")
		FORWARD_DB_PORT = viper.GetString("PG_FORWARD_DB_PORT")
		DNS_DB = viper.GetString("PG_DNS_LOCAL")

		// Variables .env HTTP_SERVER
		HTTP_SERVER_HOST = viper.GetString("HTTP_SERVER_HOST")
		HTTP_SERVER_HOST_DOC_API = viper.GetString("HTTP_SERVER_HOST_DOC_API")
		HTTP_SERVER_HOST_METRICS = viper.GetString("HTTP_SERVER_HOST_METRICS")
		HTTP_SERVER_PORT_TEST = viper.GetString("HTTP_SERVER_PORT_TEST")
		HTTP_SERVER_PORT_DOC_API = viper.GetString("HTTP_SERVER_PORT_DOC_API")
		filename := filepath.Base(os.Args[0])
		port := strings.Replace(filename, "main", "", 1)

		if port == "__debug_bin" || port == "" {
			// port = "1337"
			port = HTTP_SERVER_PORT_TEST
		} else if port != "" {
			port = strings.Replace(port, "1337", "", 1)
			port = strings.Replace(port, "133", "", 1) // JoseR
		}
		HTTP_SERVER_PORT = port
		HTTP_SERVER_PORT_METRICS = viper.GetString("HTTP_SERVER_PORT_METRICS")

		// Var env APP
		APP_MAX_CONNECTIONS = viper.GetInt("APP_MAX_CONNECTIONS")
		APP_CANT_X = viper.GetInt("APP_CANT_X")
		APP_FILE_LOGGER = viper.GetString("APP_FILE_LOGGER")
		APP_PROCESS_EVENT_RECIVED = viper.GetBool("APP_PROCESS_EVENT_RECIVED")
		APP_SESSIONS = viper.GetString("APP_SESSIONS")
		APP_AVATAR_FILES = viper.GetString("APP_AVATAR_FILES")
		APP_MESSAJE_FILES = viper.GetString("APP_MESSAJE_FILES")
		APP_NAME_X1 = viper.GetString("APP_NAME_X1")
	}
}

func prepare_app() {
	prepare_env()
	prepare_proxies()
	prepare_engine()
	// prepare_logger() //TODO:Cambiar por logecs
	// prepare_db()
	// prepare_mime_exts()

	// Test app
	//test_app()
}

func IsX1() bool {
	name := path.Base(os.Args[0])
	if name == APP_NAME_X1 || name == "main" {
		return true
	}
	return false
}

func RouterList(g *gin.Engine) {
	fmt.Printf("\n\n%d Rutas listas en %s:%s \n\n", len(g.Routes()), HTTP_SERVER_HOST, HTTP_SERVER_PORT)
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

func AppRun() {
	//RouterList(engine)
	//RouterList(metricEngine)
	httpRun()
}
