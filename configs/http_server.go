package configs

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

var (
	HTTP_SERVER_HOST         string
	HTTP_SERVER_PORT         string
	HTTP_SERVER_PORT_TEST    string
	HTTP_SERVER_HOST_DOC_API string
	HTTP_SERVER_PORT_DOC_API string

	StateInitDocApi bool = true // state for controller if server doc api is initialized or not
	StateInitMetric bool = true // state for controller if server metric is initialized or not

	HTTP_SERVER_HOST_METRICS string
	HTTP_SERVER_PORT_METRICS string
	GROUP_WAIT               errgroup.Group
)

func httpRun() {
	if IsX1() {
		go runServMetric()
		go runServerDocApi()
	}
	if err := GROUP_WAIT.Wait(); err != nil {
		fmt.Println("Ocurrio un error con la sincronizacion de server: ", err.Error())
	}
	// Servicio
	Log.Infof("Run server Api in %s:%s\n", HTTP_SERVER_HOST, HTTP_SERVER_PORT)
	err := routerApi.Run(HTTP_SERVER_HOST + ":" + HTTP_SERVER_PORT)
	if err != nil {
		Log.Errorf("Error initializing server Api error is: %s", err.Error())
	}
}

func runServMetric() {
	if HTTP_SERVER_HOST_METRICS == "" && HTTP_SERVER_PORT_METRICS == "" {
		StateInitMetric = false
		return
	}
	StateInitMetric = true
	Log.Infof("Run server Metrics in %s:%s\n", HTTP_SERVER_HOST_METRICS, HTTP_SERVER_PORT_METRICS)
	// err := metricEngine.Run(HTTP_SERVER_HOST_METRICS + ":" + HTTP_SERVER_PORT_METRICS)
	// if err != nil {
	// 	Log.Errorf("Error initializing server Metrics error is: %s", err.Error())
	// }
}

func runServerDocApi() {
	if HTTP_SERVER_HOST_DOC_API == "" && HTTP_SERVER_PORT_DOC_API == "" {
		StateInitDocApi = false
		return
	}
	StateInitDocApi = true
	Log.Infof("Run server Doc Api in %s:%s\n", HTTP_SERVER_HOST_DOC_API, HTTP_SERVER_PORT_DOC_API)
	err := routerDocApi.Run(HTTP_SERVER_HOST_DOC_API + ":" + HTTP_SERVER_PORT_DOC_API)
	if err != nil {
		Log.Errorf("Error initializing server Doc Api error is: %s", err.Error())
	}
}
