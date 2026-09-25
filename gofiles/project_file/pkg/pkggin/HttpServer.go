package pkggin

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	c_ "ecs_govel/configs"
	"ecs_govel/pkg/pkglog"

	"golang.org/x/sync/errgroup"
)

func CreateProtoHTTP2NotTLS() *http.Protocols {
	p := new(http.Protocols)
	p.SetHTTP1(true)
	p.SetUnencryptedHTTP2(true)

	return p
}

func servMetric(ctx context.Context) error {
	if c_.HTTP_SERVER_HOST_METRICS == "" && c_.HTTP_SERVER_PORT_METRICS == "" {
		return nil
	} else {
		// if c_.HTTP_SERVER_HOST_METRICS == "" {
		// 	c_.HTTP_SERVER_HOST_METRICS = c_.HTTP_SERVER_HOST
		// }
		if c_.HTTP_SERVER_PORT_METRICS == "" {
			c_.HTTP_SERVER_PORT_METRICS = "8081"
		}
	}

	host := fmt.Sprintf(":%s", c_.HTTP_SERVER_PORT_METRICS)
	serverMetric := &http.Server{
		Addr:      host,
		Handler:   metricEngine,
		Protocols: CreateProtoHTTP2NotTLS(),
	}

	go func() {
		<-ctx.Done()
		serverMetric.Shutdown(ctx)
	}()

	err := serverMetric.ListenAndServe()
	if err != nil {
		return fmt.Errorf("server Metric - %s", err.Error())
	}

	return nil
}

func servDocApi(ctx context.Context) error {
	if c_.HTTP_SERVER_HOST_DOC_API == "" && c_.HTTP_SERVER_PORT_DOC_API == "" {
		return nil
	} else {
		if c_.HTTP_SERVER_HOST_DOC_API == "" {
			c_.HTTP_SERVER_HOST_DOC_API = c_.HTTP_SERVER_HOST
		}
		if c_.HTTP_SERVER_PORT_DOC_API == "" {
			c_.HTTP_SERVER_PORT_DOC_API = "8083"
		}
	}

	host := fmt.Sprintf("%s:%s", c_.HTTP_SERVER_HOST_DOC_API, c_.HTTP_SERVER_PORT_DOC_API)

	ServerDocsApi := &http.Server{
		Addr:      host,
		Handler:   docApiEngine,
		Protocols: CreateProtoHTTP2NotTLS(),
	}

	go func() {
		<-ctx.Done()
		ServerDocsApi.Shutdown(ctx)
	}()

	err := ServerDocsApi.ListenAndServe()
	if err != nil {
		return fmt.Errorf("server DocApi - %s", err.Error())
	}

	return nil
}

func HttpRun() {
	// ServerMetrics
	GROUP_WAIT, ctx := errgroup.WithContext(context.Background()) // WithContext.
	GROUP_WAIT.SetLimit(3)
	if c_.StateInitMetric { //metricEngine
		GROUP_WAIT.Go(func() error {
			pkglog.Log.Sub("Configs").Infof("Server Metrics in %s:%s\n", c_.HTTP_SERVER_PORT, c_.HTTP_SERVER_PORT_METRICS)
			return servMetric(ctx)
		})
	}

	if c_.StateInitDocApi {
		// ServerDocsApi
		GROUP_WAIT.Go(func() error {
			pkglog.Log.Sub("Configs").Infof("Service DocsApi in %s:%s/docs\n", c_.HTTP_SERVER_HOST_DOC_API, c_.HTTP_SERVER_PORT_DOC_API)
			return servDocApi(ctx)
		})
	}

	// Servicio
	if strings.ToLower(c_.TYPE_SERVICES) == "rest" {
		host := fmt.Sprintf("%s:%s", c_.HTTP_SERVER_HOST, c_.HTTP_SERVER_PORT)

		server := &http.Server{
			Addr:      host,
			Handler:   engine,
			Protocols: CreateProtoHTTP2NotTLS(),
		}
		// http2.ConfigureServer(server, &http2.Server{})

		GROUP_WAIT.Go(func() error {
			pkglog.Log.Sub("Configs").Infof("Service Web in: %s:%s\n", c_.HTTP_SERVER_HOST, c_.HTTP_SERVER_PORT)
			go func() {
				<-ctx.Done()
				server.Shutdown(ctx)
			}()

			if err := server.ListenAndServe(); err != nil {
				return fmt.Errorf("server Web - %s", err.Error())
			}

			return nil
		})
	}

	if err := GROUP_WAIT.Wait(); err != nil {
		pkglog.Log.Sub("Configs").Errorf("Ocurr one error sync: %s\n", err.Error())
		os.Exit(2)
	}
}
