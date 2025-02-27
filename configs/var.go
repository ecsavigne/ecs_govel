package configs

import (
	logecs "github.com/ecsavigne/logecs/log"
	"github.com/ecsavigne/proxy-reverse/proxy"
)

// log
var (
	Log = logecs.NewLoggerEcs(logecs.EcsLogger{
		Mod: logAPi, Color: true,
		Path: "output.log", OutPut: true,
	})
)

// ReverseProxyApiDoc
var (
	ReverseProxyApiDoc *proxy.ProxyReverse
)
