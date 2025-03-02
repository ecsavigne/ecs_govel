package routes

import (
	c_ "ecs_govel/configs"
)

func init() {
	loadApi(c_.GetEngine())
	if c_.IsX1() {
		// loadMetrics(c_.GetMetricEngine())
		loadDocRoutes(c_.GetDocApiEngine())
	}
}
