package routes

import (
	c_ "ecs_govel/configs"
)

func init() {
	// webhookRoutes(c_.GetWebHookEngine())
	loadApi(c_.GetEngine())
	if c_.StateInitMetric && c_.IsX1() {
		loadMetrics(c_.GetMetricEngine())
	}

	if c_.StateInitDocApi && c_.IsX1() {
		e := c_.GetDocApiEngine()
		loadDocRoutes(e)
	}
}
