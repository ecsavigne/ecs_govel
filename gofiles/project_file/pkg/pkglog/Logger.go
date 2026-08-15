package pkglog

import (
	"ecs_govel/command"
	c_ "ecs_govel/configs"

	logecs "github.com/ecsavigne/logecs/log"
)

var (
	Log logecs.Logger
)

func init() {
	Log = logecs.NewLoggerEcs(logecs.EcsLogger{
		Mod: c_.LOG_APi, Color: true,
		Path: c_.APP_FILE_LOGGER, OutPut: true,
	})

	// set logger
	command.SetLogger(Log)

	Log.Infof("\x1b[34mConfiguring logger file in %s\x1b[0m\n", c_.APP_FILE_LOGGER)
}
