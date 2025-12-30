package configs

import (
	"github.com/ecsavigne/proxy-reverse/proxy"
	"github.com/gin-gonic/gin"
)

// ReverseProxyApiDoc
var (
	ReverseProxyApi *proxy.ProxyReverse // = &proxy.ProxyReverse{}
)

// App
var (
	engine                    *gin.Engine
	metricEngine              *gin.Engine
	docApiEngine              *gin.Engine
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
	DOC_API_PATH              string
	GRPC_SERVER_PORT          string
	TYPE_SERVICES             string
)
