package interceptor

import (
	"path"
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
		dirUrl, pathKey := path.Split(g.Request.URL.Path)

		if newPath, ok := routeMap[pathKey]; ok {
			param_user_id := path.Base(dirUrl)
			newPath = strings.ReplaceAll(newPath, "{ig_user_id}", param_user_id)

			g.Request.URL.Path = newPath
		} else {
			g.AbortWithStatus(404)
		}

		g.Next()
	}
}
