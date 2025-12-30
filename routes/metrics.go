package routes

import (
	"ecs_govel/configs"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func loadMetrics(g *gin.Engine) {
	g.GET("/metrics", gin.WrapH(promhttp.HandlerFor(configs.Register, promhttp.HandlerOpts{})))
}
