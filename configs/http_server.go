package configs

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

var (
	HTTP_SERVER_HOST string
	HTTP_SERVER_PORT string

	HTTP_SERVER_HOST_METRICS string
	HTTP_SERVER_PORT_METRICS string
	GROUP_WAIT               errgroup.Group
)

func httpRun() {
	// ServerMetrics
	if IsX1() { //metricEngine
		//fmt.Println("Server Metrics:" + HTTP_SERVER_HOST_METRICS + ":" + HTTP_SERVER_PORT_METRICS)
		go servMetric()
	}
	if err := GROUP_WAIT.Wait(); err != nil {
		fmt.Println("Ocurrio un error con la sincronizacion de server: ", err.Error())
	}
	// Servicio
	engine.Run(HTTP_SERVER_HOST + ":" + HTTP_SERVER_PORT)
}

func servMetric() {
	// metricEngine.Run(HTTP_SERVER_HOST_METRICS + ":" + HTTP_SERVER_PORT_METRICS)
}
