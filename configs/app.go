package configs

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	engine                    *gin.Engine
	metricEngine              *gin.Engine
	APP_MAX_CONNECTIONS       int
	APP_CANT_X                int
	APP_FILE_LOGGER           string
	APP_SESSIONS              string
	APP_MESSAJE_FILES         string
	APP_AVATAR_FILES          string
	APP_NAME_X1               string
	APP_PROCESS_EVENT_RECIVED bool
	// APIExts : Api -> wpp ext to mime convert
	APIExts                 = make(map[string]string, 40)
	DocHandlerExts          = make(map[string]string, 40)
	ImageType      []string = []string{"png", "jpg", "jpeg", "webp", "gif"}
	AudioType      []string = []string{"ogg", "mp3"}
	VideoType      []string = []string{"mp4"}
	DocumentType   []string = []string{"docx", "doc", "pdf", "epub", "ppt",
		"pptx", "csv", "xls", "xlsx"}
)

func init() {
	prepare_app()
}

func GetEngine() *gin.Engine {
	return engine
}

func GetMetricEngine() *gin.Engine {
	return metricEngine
}

// Configurar el motor de Gin
func prepare_engine() {
	gin.SetMode(gin.ReleaseMode)
	if IsX1() {
		//TODO: Configurar el motor de Gin para metricas no esta implementada la logica aun
		metricEngine = gin.Default()
	} else {
		engine = gin.Default()
	}
}

// Carga de varEnv
func prepare_env() {
	pathDir, _ := os.Getwd()
	viper.AddConfigPath(pathDir)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
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

		// Variables .env HTTP_SERVER
		HTTP_SERVER_HOST = viper.GetString("HTTP_SERVER_HOST")
		HTTP_SERVER_HOST_METRICS = viper.GetString("HTTP_SERVER_HOST_METRICS")
		filename := filepath.Base(os.Args[0])
		port := strings.Replace(filename, "main", "", 1)

		if port == "__debug_bin" || port == "" {
			port = "1337"
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
	prepare_engine()
	prepare_logger()
	prepare_db()
	// prepare_mime_exts()

	// Test app
	//test_app()
}

func test_app() {
	Database.CreateApp()
	// Database.SaveCompanyWhatsapp(models.CompanyWhatsapps{
	// 	CompanyId:            1,
	// 	CompanyWhatsappId:    1,
	// 	Whatsapp:             "5521982641843",
	// 	ProcessGroupMessages: true,
	// })
}

/*func prepare_mime_exts() {
	DocHandlerExts["application/vnd.ms-excel"] = "xls"
	DocHandlerExts["application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"] = "xlsx"
	DocHandlerExts["application/vnd.openxmlformats-officedocument.wordprocessingml.document"] = "docx"
	DocHandlerExts["application/msword"] = "doc"
	DocHandlerExts["application/vnd.ms-powerpoint"] = "ppt"
	DocHandlerExts["application/vnd.openxmlformats-officedocument.presentationml.presentation"] = "pptx"

	DocHandlerExts["application/vnd.oasis.opendocument.text"] = "odt"
	DocHandlerExts["application/vnd.oasis.opendocument.spreadsheet"] = "ods"
	DocHandlerExts["application/vnd.oasis.opendocument.presentation"] = "odp"

	DocHandlerExts["application/epub+zip"] = "epub"
	DocHandlerExts["text/plain"] = "text"

	DocHandlerExts["text/csv"] = "csv"

	DocHandlerExts["video/mp4"] = "mp4"

	DocHandlerExts["application/pdf"] = "pdf"
	DocHandlerExts["application/zip"] = "zip"

	DocHandlerExts["image/png"] = "png"
	DocHandlerExts["image/jpg"] = "jpg"
	DocHandlerExts["image/jpeg"] = "jpeg"
	DocHandlerExts["image/gif"] = "gif"

	DocHandlerExts["video/mpeg"] = "mpeg"
	DocHandlerExts["image/webp"] = "webp"

	DocHandlerExts["audio/mpeg"] = "mp3"
	// DocHandlerExts["audio/ogg"] = "ogg"
	DocHandlerExts["audio/ogg; codecs=opus"] = "ogg"
	DocHandlerExts["audio/mpeg; codecs=opus"] = "mpeg"
	DocHandlerExts["audio/mp4; codecs=opus"] = "mp4"
	DocHandlerExts["audio/aac; codecs=opus"] = "aac"

	//DocHandlerExts[""] = "rem"
	DocHandlerExts["application/ofx"] = "ofx"

	APIExts["png"] = "image/png"
	APIExts["jpg"] = "image/jpeg"
	APIExts["jpeg"] = "image/jpeg"
	// added by JR
	APIExts["webp"] = "image/webp"

	APIExts["docx"] = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	APIExts["doc"] = "application/msword"
	APIExts["pdf"] = "application/pdf"
	APIExts["epub"] = "application/epub+zip"

	APIExts["ppt"] = "application/vnd.ms-powerpoint"
	APIExts["pptx"] = "application/vnd.openxmlformats-officedocument.presentationml.presentation"

	APIExts["gif"] = "image/gif"
	APIExts["csv"] = "application/csv" // APIExts["csv"] = "text/csv"
	APIExts["xlsx"] = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	APIExts["xls"] = "application/excel"

	// video format
	APIExts["mp4"] = "video/mp4"

	// audio formats
	APIExts["ogg"] = "audio/ogg; codecs=opus"
	APIExts["acc"] = "audio/acc; codecs=opus"
	APIExts["aac"] = "audio/ogg; codecs=opus"
	APIExts["mp3"] = "audio/mpeg"

	APIExts["txt"] = "text/plain"
	APIExts["odt"] = "application/vnd.oasis.opendocument.text"
	APIExts["zip"] = "application/zip"

	// APIExts["rem"] = "text/plain"
	// APIExts["ofx"] = "application/ofx"

}*/

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
