package restroute

import (
	"ecs_govel/pkg/pkgmetrics/sdkopentelemetry"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func loadMetrics(g *gin.Engine) {
	g.GET("/metrics", gin.WrapH(promhttp.HandlerFor(sdkopentelemetry.Register, promhttp.HandlerOpts{})))
}
