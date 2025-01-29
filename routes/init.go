package routes

import (
	c_ "ecs_govel/configs"
)

func init() {
	loadApi(c_.GetEngine())
	// loadMetrics(c_.GetMetricEngine())
}
