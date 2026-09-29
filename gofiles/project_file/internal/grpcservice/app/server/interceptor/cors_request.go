package interceptor

import (
	"ecs_govel/pkg/pkggin"
	"slices"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsReq() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			origin = strings.TrimPrefix(strings.Split(origin, ":")[1], "//")
			_, ok := slices.BinarySearch(pkggin.HostAllow(), origin)
			return ok
		},
		AllowMethods: []string{"GET", "POST", "OPTIONS", "DELETE", "PUT", "PATCH"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
			"Connect-Protocol-Version",
			"Connect-Timeout-Ms",
		},
		MaxAge: 12 * time.Hour,
	})
}
