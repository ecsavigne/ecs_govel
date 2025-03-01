package configs

import logecs "github.com/ecsavigne/logecs/log"

// log
var (
	Log logecs.Logger
)

func prepare_logger() {
	Log = logecs.NewLoggerEcs(logecs.EcsLogger{
		Mod: logAPi, Color: true,
		Path: APP_FILE_LOGGER, OutPut: true,
	})
}
