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

func GetDocApiEngine() *gin.Engine {
	return docApiEngine
}

// Configurar el motor de Gin
func prepare_engine() {
	if strings.ToLower(APP_MODE) == "develop" {
		fmt.Println("dEVELOP MODE")
		gin.SetMode(gin.DebugMode)
	} else {
		fmt.Println("rELEASE MODE")
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.MultiWriter(fileLogger, os.Stdout)
	}

	if IsX1() {
		//TODO: Configurar el motor de Gin para metricas no esta implementada la logica aun
		// fmt.Println("Configurar el motor de Gin para metricas y documentacion")
		Log.Sub("Configs").Infof("Configurar el motor de Gin para metricas y documentacion\n")
		docApiEngine = gin.Default()

		metricEngine = gin.Default()
	}

	engine = gin.Default()
	engine.MaxMultipartMemory = 100 << 20
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
		DB_TYPE = viper.GetString("DB_TYPE")
		postgres := func() {
			PG_DB_HOST = viper.GetString("PG_DB_HOST")
			PG_DB_USER = viper.GetString("PG_DB_USERNAME")
			PG_DB_NAME = viper.GetString("PG_DB_DATABASE")
			PG_DB_PASSWORD = viper.GetString("PG_DB_PASSWORD")
			PG_DB_PORT = viper.GetString("PG_DB_PORT")
			PG_DNS_DB = viper.GetString("PG_DNS_LOCAL")
			DB_TYPE = "postgres"
		}

		mongo := func() {
			MONGO_DB_HOST = viper.GetString("MONGO_DB_HOST")
			MONGO_DB_USER = viper.GetString("MONGO_DB_USERNAME")
			MONGO_DB_NAME = viper.GetString("MONGO_DB_DATABASE")
			MONGO_DB_PASSWORD = viper.GetString("MONGO_DB_PASSWORD")
			MONGO_DB_PORT = viper.GetString("MONGO_DB_PORT")
		}

		switch DB_TYPE {
		case "":
			fallthrough
		case "postgres":
			postgres()
		case "mongo":
			mongo()
		case "all":
			postgres()
			mongo()
		}
		FORWARD_DB_PORT = viper.GetString("PG_FORWARD_DB_PORT")

		SSH_ENABLE, err = strconv.ParseBool(viper.GetString("SSH_ENABLE"))
		if err != nil {
			SSH_ENABLE = false
		}
		SSH_PORT = viper.GetString("SSH_PORT")
		SSH_HOST = viper.GetString("SSH_HOST")
		SSH_PASS = viper.GetString("SSH_PASS")
		SSH_USER = viper.GetString("SSH_USER")

		// Var webhook
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
		DOC_API_PATH = viper.GetString("DOC_API_PATH")

		// Variables .env HTTP_SERVER
		HTTP_SERVER_HOST = viper.GetString("HTTP_SERVER_HOST")
		HTTP_SERVER_HOST_WEBHOOK = viper.GetString("HTTP_SERVER_HOST_WEBHOOK")
		HTTP_SERVER_HOST_METRICS = viper.GetString("HTTP_SERVER_HOST_METRICS")
		HTTP_SERVER_HOST_DOC_API = viper.GetString("HTTP_SERVER_HOST_DOC_API")
		HTTP_SERVER_PORT_TEST = viper.GetString("HTTP_SERVER_PORT_TEST")
		HTTP_SERVER_PORT_DOC_API = viper.GetString("HTTP_SERVER_PORT_DOC_API")
		HTTP_SERVER_PORT_METRICS = viper.GetString("HTTP_SERVER_PORT_METRICS")
		HTTP_SERVER_PORT_WEBHOOK = viper.GetString("HTTP_SERVER_PORT_WEBHOOK")
		TYPE_SERVICES = viper.GetString("TYPE_SERVICES")
		TYPE_DOCUMENTATION = viper.GetString("TYPE_DOCUMENTATION")

		if strings.ToLower(APP_MODE) != "develop" && strings.ToLower(APP_MODE) != "production" {
			Log.Sub("configs").Errorf("APP_MODE is not valid in app.env, values possible: [develop, production] \n")
			os.Exit(2)
		}

		if TYPE_SERVICES == "" {
			Log.Sub("configs").Errorf("TYPE_SERVICES is empty in app.env, values possible: [grpc, rest] \n")
			os.Exit(2)
		} else {
			switch strings.ToLower(TYPE_SERVICES) {
			case "grpc":
				GRPC_SERVER_PORT = viper.GetString("GRPC_SERVER_PORT")
				if GRPC_SERVER_PORT == "" {
					Log.Sub("configs").Errorf("GRPC_SERVER_PORT is empty in app.env\n")
					os.Exit(2)
				}

				if _, e := strconv.Atoi(GRPC_SERVER_PORT); e != nil {
					Log.Sub("configs").Errorf("GRPC_SERVER_PORT is not valid in app.env: error is: %s\n", e.Error())
					os.Exit(2)
				}

			case "rest":
				HTTP_SERVER_PORT = "8080"
				if strings.ToLower(APP_MODE) == "develop" {
					if HTTP_SERVER_PORT_TEST != "" {
						HTTP_SERVER_PORT = HTTP_SERVER_PORT_TEST
					}
				} else if strings.ToLower(APP_MODE) == "production" {
					HTTP_SERVER_PORT = strings.TrimPrefix(path.Base(os.Args[0]), APP_NAME)
					if _, e := strconv.Atoi(HTTP_SERVER_PORT); e != nil {
						HTTP_SERVER_PORT = "8080"
					}
				}
				Log.Sub("configs").Debugf("prepare_env HTTP_SERVER_Port: %s\n", HTTP_SERVER_PORT)
			default:
				Log.Sub("configs").Errorf("TYPE_SERVICES is not valid view in app.env, values possible: [grpc, rest]. Actual value: %s\n", strings.ToLower(TYPE_SERVICES))
				os.Exit(2)
			}
		}

		if HTTP_SERVER_HOST_METRICS == "" && HTTP_SERVER_PORT_METRICS == "" {
			StateInitMetric = false
		}

	}
}

func prepare_app() {
	fmt.Println("prepare_app")
	prepare_env()
	prepare_engine()
	prepare_interceptor()
	// prepare_db()
	// prepare_mime_exts()
}

func IsX1() bool {
	name := path.Base(os.Args[0])
	if name == APP_NAME_X1 || strings.ToLower(APP_MODE) == "develop" || "ecs_govel" == name {
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
