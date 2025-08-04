package configs

import (
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	engine                    *gin.Engine
	metricEngine              *gin.Engine
	webHookEngine             *gin.Engine
	APP_MAX_CONNECTIONS       int
	APP_CANT_X                int
	APP_FILE_LOGGER           string
	APP_URL_BASE_WEBHOOK      string
	APP_URL_FILE_STORE        string
	APP_SESSIONS              string
	APP_MESSAJE_FILES         string
	APP_AVATAR_FILES          string
	APP_NAME                  string
	APP_NAME_X1               string
	APP_PROCESS_EVENT_RECIVED bool
	APP_MODE                  string
	APP_PORT_TEST             string
	WEBHOOK_SOCKET            string
)

func init() {
	fmt.Println("init Config package")
	prepare_app()
}

func GetEngine() *gin.Engine {
	return engine
}

func GetMetricEngine() *gin.Engine {
	return metricEngine
}

func GetWebHookEngine() *gin.Engine {
	return webHookEngine
}

// Configurar el motor de Gin
func prepare_engine() {
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	gin.SetMode(gin.ReleaseMode)
	if IsX1() {
		//TODO: Configurar el motor de Gin para metricas no esta implementada la logica aun
		metricEngine = gin.Default()
	}

	engine = gin.Default()
	engine.MaxMultipartMemory = 100 << 20

	webHookEngine = gin.Default()
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
		WEBHOOK_SOCKET = viper.GetString("WEBHOOK_SOCKET")

		// Var env APP
		APP_MODE = viper.GetString("APP_MODE")
		APP_PORT_TEST = viper.GetString("APP_PORT_TEST")
		APP_MAX_CONNECTIONS = viper.GetInt("APP_MAX_CONNECTIONS")
		APP_CANT_X = viper.GetInt("APP_CANT_X")
		APP_FILE_LOGGER = viper.GetString("APP_FILE_LOGGER")
		APP_URL_FILE_STORE = viper.GetString("APP_URL_FILE_STORE")
		APP_URL_BASE_WEBHOOK = viper.GetString("APP_URL_BASE_WEBHOOK")
		prepare_logger()
		APP_PROCESS_EVENT_RECIVED = viper.GetBool("APP_PROCESS_EVENT_RECIVED")
		APP_SESSIONS = viper.GetString("APP_SESSIONS")
		APP_AVATAR_FILES = viper.GetString("APP_AVATAR_FILES")
		APP_MESSAJE_FILES = viper.GetString("APP_MESSAJE_FILES")
		APP_NAME_X1 = viper.GetString("APP_NAME_X1")
		APP_NAME = viper.GetString("APP_NAME")
		APP_MODE = viper.GetString("APP_MODE")

		// Variables .env HTTP_SERVER
		HTTP_SERVER_HOST = viper.GetString("HTTP_SERVER_HOST")
		HTTP_SERVER_HOST_WEBHOOK = viper.GetString("HTTP_SERVER_HOST_WEBHOOK")
		HTTP_SERVER_HOST_METRICS = viper.GetString("HTTP_SERVER_HOST_METRICS")
		HTTP_SERVER_HOST_DOC_API = viper.GetString("HTTP_SERVER_HOST_DOC_API")
		HTTP_SERVER_PORT_TEST = viper.GetString("HTTP_SERVER_PORT_TEST")
		HTTP_SERVER_PORT_DOC_API = viper.GetString("HTTP_SERVER_PORT_DOC_API")
		HTTP_SERVER_PORT_METRICS = viper.GetString("HTTP_SERVER_PORT_METRICS")
		HTTP_SERVER_PORT_WEBHOOK = viper.GetString("HTTP_SERVER_PORT_WEBHOOK")

		port := "8080"
		if strings.ToLower(APP_MODE) == "develop" {
			if HTTP_SERVER_PORT_TEST != "" {
				port = HTTP_SERVER_PORT_TEST
			}
		} else if strings.ToLower(APP_MODE) == "production" {
			port = strings.TrimPrefix(path.Base(os.Args[0]), APP_NAME)
			if _, e := strconv.Atoi(port); e != nil {
				port = "8080"
			}
		}
		Log.Debugf("prepare_env HTTP_SERVER_Port: %s\n", port)

		HTTP_SERVER_PORT = port
		if HTTP_SERVER_HOST_METRICS == "" && HTTP_SERVER_PORT_METRICS == "" {
			StateInitMetric = false
		}

	}
}

func prepare_app() {
	fmt.Println("prepare_app")
	prepare_env()
	prepare_engine()
	prepare_db()
	// prepare_mime_exts()
}

func IsX1() bool {
	name := path.Base(os.Args[0])
	if name == APP_NAME_X1 || strings.ToLower(APP_MODE) == "develop" {
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
