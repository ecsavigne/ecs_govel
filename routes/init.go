package routes

import (
	c_ "ecs_govel/configs"
)

func init() {
	loadApi(c_.GetEngine())
	if c_.StateInitDocApi {
		// loadMetrics(c_.GetMetricEngine())
		loadDocRoutes(c_.GetDocApiEngine())
	}
}
