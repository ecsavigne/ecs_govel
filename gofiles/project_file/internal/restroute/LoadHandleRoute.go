package restroute

import (
	c_ "ecs_govel/configs"
	"ecs_govel/pkg/pkggin"
)

func loadHandlerApiRoute() {
	// webhookRoutes(pkggin.GetWebHookEngine())
	loadApi(pkggin.GetEngine())
}

func loadHandlerDocRoute() {
	if c_.StateInitDocApi && c_.IsX1() {
		e := pkggin.GetDocApiEngine()
		loadDocRoutes(e)
	}
}

func loadHandlerMetricsRoute() {
	if c_.StateInitMetric && c_.IsX1() {
		loadMetrics(pkggin.GetMetricEngine())
	}
}

func LoadAllHandlers() {
	loadHandlerApiRoute()
	loadHandlerDocRoute()
	loadHandlerMetricsRoute()
}

func LoadInfoHandlers() {
	loadHandlerDocRoute()
	loadHandlerMetricsRoute()
}
