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

	if c_.StateInitDocApi {
		e := c_.GetDocApiEngine()
		c_.Log.Infof("loadDocRoutes is: %v\n\n", e)
		loadDocRoutes(e)
	}
}
