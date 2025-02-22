package configs

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

var (
	HTTP_SERVER_HOST         string
	HTTP_SERVER_PORT         string
	HTTP_SERVER_PORT_TEST    string
	HTTP_SERVER_PORT_DOC_API string

	HTTP_SERVER_HOST_METRICS string
	HTTP_SERVER_PORT_METRICS string
	GROUP_WAIT               errgroup.Group
)

func httpRun() {
	// ServerMetrics
	if IsX1() { //metricEngine
		//fmt.Println("Server Metrics:" + HTTP_SERVER_HOST_METRICS + ":" + HTTP_SERVER_PORT_METRICS)
		go runServMetric()
	}
	if err := GROUP_WAIT.Wait(); err != nil {
		fmt.Println("Ocurrio un error con la sincronizacion de server: ", err.Error())
	}
	// Servicio
	fmt.Printf("Run server Api in %s:%s\n", HTTP_SERVER_HOST, HTTP_SERVER_PORT)
	routerApi.Run(HTTP_SERVER_HOST + ":" + HTTP_SERVER_PORT)
}

func runServMetric() {
	fmt.Printf("Run server Metrics in %s:%s\n", HTTP_SERVER_HOST_METRICS, HTTP_SERVER_PORT_METRICS)
	// metricEngine.Run(HTTP_SERVER_HOST_METRICS + ":" + HTTP_SERVER_PORT_METRICS)
}

func runServerDocApi() {
	fmt.Printf("Run server Doc Api in %s:%s\n", HTTP_SERVER_HOST, HTTP_SERVER_PORT_DOC_API)
	routerDocApi.Run(HTTP_SERVER_HOST + ":" + HTTP_SERVER_PORT_DOC_API)
}
