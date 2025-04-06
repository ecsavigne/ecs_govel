package configs

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
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
		GROUP_WAIT.Go(func() error {
			Log.Infof("Run server Metrics in %s:%s\n", HTTP_SERVER_HOST_METRICS, HTTP_SERVER_PORT_METRICS)
			return runServMetric()
		})

		GROUP_WAIT.Go(func() error {
			Log.Infof("Run server Doc Api in %s:%s\n", HTTP_SERVER_HOST_DOC_API, HTTP_SERVER_PORT_DOC_API)
			return runServerDocApi()
		})
	}

	// Servicio
	server := &http.Server{
		Addr:    HTTP_SERVER_HOST + ":" + HTTP_SERVER_PORT,
		Handler: routerApi,
	}
	err := http2.ConfigureServer(server, &http2.Server{})
	if err != nil {
		Log.Errorf("Error configuring server Api error is: %s", err.Error())
	}

	GROUP_WAIT.Go(func() error {
		Log.Infof("Run server Api in %s:%s", HTTP_SERVER_HOST, HTTP_SERVER_PORT)
		return server.ListenAndServe()
	})
	server.ListenAndServe()

	if err := GROUP_WAIT.Wait(); err != nil {
		fmt.Printf("Ocurred one error with the sincronization of server error is: %s\n", err.Error())
	}
}

func runServMetric() error {
	if HTTP_SERVER_HOST_METRICS == "" && HTTP_SERVER_PORT_METRICS == "" {
		return nil
	} else {
		if HTTP_SERVER_HOST_METRICS == "" {
			HTTP_SERVER_HOST_METRICS = HTTP_SERVER_HOST
		}
		if HTTP_SERVER_PORT_METRICS == "" {
			HTTP_SERVER_PORT_METRICS = "8081"
		}
	}

	// metricServer := &http.Server{
	// 	Addr:    HTTP_SERVER_HOST_METRICS + ":" + HTTP_SERVER_PORT_METRICS,
	// 	Handler: routerMetric,
	// }

	// err := http2.ConfigureServer(metricServer, &http2.Server{})
	// if err != nil {
	// 	Log.Errorf("Error configuring server Metrics error is: %s", err.Error())
	// }

	// go func() {
	// err = metricServer.ListenAndServe()
	// if err != nil {
	// 	Log.Errorf("Error initializing server Metrics error is: %s", err.Error())
	// }
	// }
	return nil
}

func runServerDocApi() error {
	if HTTP_SERVER_HOST_DOC_API == "" && HTTP_SERVER_PORT_DOC_API == "" {
		return nil
	} else {
		if HTTP_SERVER_HOST_DOC_API == "" {
			HTTP_SERVER_HOST_DOC_API = HTTP_SERVER_HOST
		}
		if HTTP_SERVER_PORT_DOC_API == "" {
			HTTP_SERVER_PORT_DOC_API = "8082"
		}
	}

	// serverDocApi := &http.Server{
	// 	Addr:    HTTP_SERVER_HOST_DOC_API + ":" + HTTP_SERVER_PORT_DOC_API,
	// 	Handler: routerDocApi,
	// }

	// err := http2.ConfigureServer(serverDocApi, &http2.Server{})
	// if err != nil {
	// 	Log.Errorf("Error configuring server Doc Api error is: %s", err.Error())
	// }

	// go func() {
	// err = serverDocApi.ListenAndServe()
	// if err != nil {
	// 	Log.Errorf("Error initializing server Doc Api error is: %s", err.Error())
	// }
	// }()
	return nil
}
