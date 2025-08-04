package routes

import (
	c_ "ecs_govel/configs"
)

func init() {
	// webhookRoutes(c_.GetWebHookEngine())
	loadApi(c_.GetEngine())
	if c_.StateInitMetric {
		// loadMetrics(c_.GetMetricEngine())
	}
}
