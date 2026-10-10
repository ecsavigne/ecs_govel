package pkgutil

import (
	"ecs_govel/pkg/pkglog"
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"

	logecs "github.com/ecsavigne/logecs/log"
)

func ParseHost(host string) string {
	host, _, _ = strings.Cut(strings.NewReplacer("https://", "", "http://", "").Replace(host), ":")

	return host
}

func TryGRPC(err_return *error, debugs ...*string) {
	if r := recover(); r != nil {
		d := ""
		if len(debugs) > 0 {
			d = *debugs[0]
		}

		panicStr := ""

		if e, ok := r.(runtime.Error); ok {
			panicStr = logecs.Str(string(debug.Stack())).Red().String()
			pkglog.Log.Create(logecs.InfoLog{Type: logecs.Error, Sub: "configs.TryGRPC", Name: "TryGRPC", Content: map[string]any{"action": "panic", "recovery_value": logecs.Str(e.Error()).Yellow().Bold().Italics().String(), "stack": panicStr}}) //e.Error()e.Error(), "stack": panicStr}})

			*err_return = fmt.Errorf("recovery from panic: Debug info is: %+v,  panic is: %+v", d, panicStr)

			return
		}

		*err_return = fmt.Errorf("recovery from panic: Debug info is: %+v, recovery value is: %+v", d, r)
		return
	}
}
