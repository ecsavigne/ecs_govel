package routes

import (
	c_ "oficial_gin/configs"
)

func init() {
	loadApi(c_.GetEngine())
	// loadMetrics(c_.GetMetricEngine())
}
