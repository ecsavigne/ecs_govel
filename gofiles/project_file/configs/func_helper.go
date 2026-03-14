//lint:file-ignore ST1005 Ignore capitalized strings error
package configs

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Try recovers from panics and returns a 200 json response with the message of the panic.
// If the debug parameter is not empty, it is included in the response.
// If the http.ResponseWriter is nil, it logs the error and does nothing.
func Try(w http.ResponseWriter, debug ...*string) {

	if r := recover(); r != nil {
		d := ""
		if len(debug) > 0 {
			d = *debug[0]
		}

		if w == nil {
			Log.Errorf("RECOVERY FROM EXCEPTION: %+v, Debug info is: %+v\n", r, d)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"Status":    false,
			"Message":   fmt.Sprintf("RECOVERY FROM EXCEPTION: %+v. Debug info: %+v", r, d),
			"Code":      "impossible proces http request in golang server",
			"Data":      time.Now().Format("2006-01-02 15:04:05"),
			"ErrorCode": -1,
		})
	}
}

// ValidParamsNotEmpy check if all params in map[string]any are not empty,
// if some param is empty, return false and error.
// Supported types:
// - string
// - int, int8, int16, int32, int64
// - uint, uint8, uint16, uint32, uint64
// - float32, float64
func ValidParamsNotEmpy(data map[string]any) (bool, error) {
	for k, v := range data {
		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			if v.String() == "" {
				return false, fmt.Errorf("Param %s is required", k)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if v.Int() == 0 {
				return false, fmt.Errorf("Param %s is required", k)
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if v.Uint() == 0 {
				return false, fmt.Errorf("Param %s is required", k)
			}
		case reflect.Float32, reflect.Float64:
			if v.Float() == 0 {
				return false, fmt.Errorf("Param %s is required", k)
			}
		}
	}
	return true, nil
}

// CleanFormData removes all temporary files associated with the multipart form
// data in the provided gin context. This is useful for cleaning up resources
// after handling file uploads.
// Returns an error if the removal fails.
func CleanFormData(g *gin.Context) error {
	f, _ := g.MultipartForm()
	return f.RemoveAll()
}

// CreateArray splits a string `s` into a slice of substrings using the specified `separator`.
// Each substring is trimmed of leading and trailing spaces.
// Returns a slice of the trimmed substrings.
func CreateArray(s, separator string) []string {
	array := strings.Split(s, separator)

	for i, v := range array {
		array[i] = strings.TrimSuffix(strings.TrimPrefix(v, " "), " ")
	}

	return array
}

// ValidKey checks if all the keys in the list exist in the data map.
// If at least one key is not found, returns an error.
// If the list of keys is empty, returns an error.
func ValidKey(data map[string]any, key ...string) error {
	if len(key) == 0 {
		return errors.New("key is required. Need at least one key to validate")
	}

	for _, v := range key {
		if _, ok := data[v]; !ok {
			return fmt.Errorf("key %s not found", v)
		}
	}

	return nil
}
