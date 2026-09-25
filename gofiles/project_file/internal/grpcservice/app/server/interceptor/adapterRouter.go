package interceptor

import (
	"ecs_govel/configs"
	"maps"
	"net/http"
	"path"
	"path/filepath"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

var routeMap = map[string]string{
	"getHistoryPage":      "/ig_message_service/{ig_user_id}/messages/get_history_page",
	"SendAudioMessage":    "/ig_message_service/:param/messages/audio",
	"SharedContact":       "/ig_message_service/{ig_user_id}/messages/contact",
	"SendDocumentMessage": "/ig_message_service/{ig_user_id}/messages/document",
	"SendImageMessage":    "/ig_message_service/{ig_user_id}/messages/image",
	"SendLocationMessage": "/ig_message_service/{ig_user_id}/messages/location",
	"SendTextMessage":     "/ig_message_service/{ig_user_id}/messages/text",
	// "":                 "/ig_message_service/{ig_user_id}/messages/edit",
	"SendReactiontMessage": "/ig_message_service/{ig_user_id}/messages/reaction",
	"SendResponseMessage":  "/ig_message_service/{ig_user_id}/messages/response",
	"SendStickerMessage":   "/ig_message_service/{ig_user_id}/messages/sticker",
	"SendUrlMessage":       "/ig_message_service/{ig_user_id}/messages/url",
	"SendVideoMessage":     "/ig_message_service/{ig_user_id}/messages/video",
}

// Midleware que resolve routes
func AdapterMiddleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		keys := slices.Collect(maps.Keys(routeMap))

		switch {
		case slices.Contains(keys, filepath.Base(g.Request.URL.Path)):
			if b := strings.Split(g.Request.URL.Path, "/"); b[1] != configs.HASH_ROUTE {
				g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "route not authorized"})
				return
			}

			dirUrl, pathKey := filepath.Split(g.Request.URL.Path)

			if newPath, ok := routeMap[pathKey]; ok {
				param_account_id := path.Base(dirUrl)
				newPath = strings.Replace(newPath, "{ig_account_id}", param_account_id, 1)
				g.Request.URL.Path = newPath
			}
		case strings.Contains(g.Request.URL.Path, configs.HASH_ROUTE):
			g.Request.URL.Path = strings.TrimPrefix(g.Request.URL.Path, configs.HASH_ROUTE)
		default:
			g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "route not authorized"})
			return
		}

		g.Next()
	}
}
