package configs

import (
	"os"

	"golang.org/x/sync/errgroup"

	logecs "github.com/ecsavigne/logecs/log"
)

// App
var (
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
	TYPE_DOCUMENTATION        string
	PATH_BASE                 string
)

// DB

var (
	PG_DB_HOST        string
	PG_DB_USER        string
	PG_DB_NAME        string
	PG_DB_PASSWORD    string
	PG_DB_PORT        string
	FORWARD_DB_PORT   string
	PG_DNS_DB         string
	MONGO_DB_HOST     string
	MONGO_DB_USER     string
	MONGO_DB_NAME     string
	MONGO_DB_PASSWORD string
	MONGO_DB_PORT     string
	MONGO_DB_CONNSTR  string
	DB_TYPE           string
	SSH_ENABLE        bool
	SSH_PORT          string
	SSH_HOST          string
	SSH_PASS          string
	SSH_USER          string
)

// Http server
var (
	HTTP_SERVER_HOST string
	HTTP_SERVER_PORT string

	HTTP_SERVER_HOST_WEBHOOK string
	HTTP_SERVER_PORT_WEBHOOK string

	HTTP_SERVER_HOST_METRICS string
	HTTP_SERVER_PORT_METRICS string

	SERVER_PORT_TEST         string
	HTTP_SERVER_HOST_DOC_API string
	HTTP_SERVER_PORT_DOC_API string
	GROUP_WAIT               errgroup.Group

	StateInitDocApi bool = true
	StateInitMetric bool = true
)

// log
var (
	configlog  logecs.Logger
	fileLogger *os.File
)

// Telemetry
var (
	PROMETHEUS_CONFIG_PATH string
	GRAFANA_CONFIG_PATH    string
)
