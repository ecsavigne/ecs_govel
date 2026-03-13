package configs

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func secureServer(route *gin.Engine, expectHost string) {
	route.Use(func(c *gin.Context) {
		if c.Request.Host != expectHost {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
			return
		}

		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		c.Next()
	})
}

func CreateProtoHTTP2NotTLS() *http.Protocols {
	p := new(http.Protocols)
	p.SetHTTP1(true)
	p.SetUnencryptedHTTP2(true)

	return p
}

func httpRun() {
	// ServerMetrics
	GROUP_WAIT, ctx := errgroup.WithContext(context.Background()) // WithContext.
	GROUP_WAIT.SetLimit(3)
	if IsX1() { //metricEngine
		// GROUP_WAIT.Go(func() error {
		// 	Log.Sub("Configs").Infof("Server Metrics in %s:%s\n", HTTP_SERVER_HOST_METRICS, HTTP_SERVER_PORT_METRICS)
		// 	return servMetric(ctx)
		// })

		// ServerDocsApi
		GROUP_WAIT.Go(func() error {
			Log.Sub("Configs").Infof("Service DocsApi in %s:%s/docs\n", HTTP_SERVER_HOST_DOC_API, HTTP_SERVER_PORT_DOC_API)
			return servDocApi(ctx)
		})
	}

	// Servicio
	if strings.ToLower(TYPE_SERVICES) == "rest" {
		host := fmt.Sprintf("%s:%s", HTTP_SERVER_HOST, HTTP_SERVER_PORT)
		secureServer(engine, host)

		server := &http.Server{
			Addr:      host,
			Handler:   engine,
			Protocols: CreateProtoHTTP2NotTLS(),
		}
		// http2.ConfigureServer(server, &http2.Server{})

		GROUP_WAIT.Go(func() error {
			Log.Sub("Configs").Infof("Service Web in: %s:%s\n", HTTP_SERVER_HOST, HTTP_SERVER_PORT)
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

	// // ServerWebhook
	// GROUP_WAIT.Go(func() error {
	// 	fmt.Printf("Service Webhook in %s:%s\n", HTTP_SERVER_HOST_WEBHOOK, HTTP_SERVER_PORT_WEBHOOK)
	// 	return servWebhook()
	// })

	if err := GROUP_WAIT.Wait(); err != nil {
		Log.Sub("Configs").Errorf("Ocurr one error sync: %s\n", err.Error())
		os.Exit(2)
	}
}

func servMetric(ctx context.Context) error {
	if HTTP_SERVER_HOST_METRICS == "" && HTTP_SERVER_PORT_METRICS == "" {
		return nil
	} else {
		// if HTTP_SERVER_HOST_METRICS == "" {
		// 	HTTP_SERVER_HOST_METRICS = HTTP_SERVER_HOST
		// }
		if HTTP_SERVER_PORT_METRICS == "" {
			HTTP_SERVER_PORT_METRICS = "8081"
		}
	}

	host := fmt.Sprintf(":%s", HTTP_SERVER_PORT_METRICS)
	secureServer(metricEngine, host)
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

func servWebhook() error {
	// if HTTP_SERVER_HOST_WEBHOOK == "" && HTTP_SERVER_PORT_WEBHOOK == "" {
	// 	return nil
	// } else {
	// 	if HTTP_SERVER_HOST_WEBHOOK == "" {
	// 		HTTP_SERVER_HOST_WEBHOOK = HTTP_SERVER_HOST
	// 	}
	// 	if HTTP_SERVER_PORT_WEBHOOK == "" {
	// 		HTTP_SERVER_PORT_WEBHOOK = "8083"
	// 	}
	// }

	// host := fmt.Sprintf("%s:%s", HTTP_SERVER_HOST_WEBHOOK, HTTP_SERVER_PORT_WEBHOOK)
	// secureServer(metricEngine, host)
	// serverWEBHOOK := &http.Server{
	// 	Addr:    host,
	// 	Handler: docApiEngine,
	// Protocols: CreateProtoHTTP2NotTLS(),
	// }

	//Not err := http2.ConfigureServer(serverWEBHOOK, &http2.Server{})
	// if err != nil {
	// 	fmt.Println("Ocurrio un error con la configuring de server WEBHOOK: ", err.Error())
	// }

	// return serverWEBHOOK.ListenAndServe()
	return nil
}

func servDocApi(ctx context.Context) error {
	if HTTP_SERVER_HOST_DOC_API == "" && HTTP_SERVER_PORT_DOC_API == "" {
		return nil
	} else {
		if HTTP_SERVER_HOST_DOC_API == "" {
			HTTP_SERVER_HOST_DOC_API = HTTP_SERVER_HOST
		}
		if HTTP_SERVER_PORT_DOC_API == "" {
			HTTP_SERVER_PORT_DOC_API = "8083"
		}
	}

	host := fmt.Sprintf("%s:%s", HTTP_SERVER_HOST_DOC_API, HTTP_SERVER_PORT_DOC_API)

	secureServer(docApiEngine, host)

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
