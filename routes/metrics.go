package routes

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func loadMetrics(g *gin.Engine) {
	folderRoot, _ := os.Getwd()
	g.LoadHTMLGlob(folderRoot + "/template/*")

	g.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Server Metrics",
		})
	})

	// g.GET("/metrics/bd", func(c *gin.Context) {
	// 	promHandler.ServeHTTP(c.Writer, c.Request)
	// })

	/*opsProcessed := promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events 11111111111111111111",
	})

	cpuUtilization := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "colletor_tester",
		Help: "Utilización de CPU de la base de datos PostgreSQL",
	})

	reg := prometheus.NewRegistry()
	reg.MustRegister(opsProcessed, cpuUtilization)

	prometheus.NewGaugeFunc()

	var i float64 = 0
	func() {
		go func() {
			for {
				opsProcessed.Inc()
				i = i + 0.8
				cpuUtilization.Set(i)
				time.Sleep(2 * time.Second)
			}
		}()
	}()*/

	// g.GET("/metrics/trafics", gin.WrapH(promhttp.Handler()))
	g.GET("/metrics/trafics", gin.WrapH(promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{})))

	// g.GET("/metrics", gin.WrapH(promhttp.HandlerFor(reg, promhttp.HandlerOpts{
	// 	EnableOpenMetrics: true,
	// })))
}
