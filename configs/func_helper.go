package configs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func Try(w http.ResponseWriter, debug ...*string) {
	if r := recover(); r != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		d := ""
		if len(debug) > 0 {
			d = *debug[0]
		}
		_ = json.NewEncoder(w).Encode(map[string]any{
			"Status":    false,
			"Message":   fmt.Sprintf("RECOVERY FROM EXCEPTION: %+v. Debug info: %+v", r, d),
			"Code":      "impossible proces http request in golang server",
			"Data":      time.Now().Format("2006-01-02 15:04:05"),
			"ErrorCode": -1,
		})
	}
}
