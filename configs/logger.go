package configs

import (
	"os"
	"path"

	logecs "github.com/ecsavigne/logecs/log"
)

// log
var (
	Log, LogTemp logecs.Logger
)

func prepare_logger() {
	LogTemp := logecs.NewLoggerEcs(logecs.EcsLogger{
		Mod: logAPi, Color: true,
		OutPut: false,
	})

	folderLog := path.Dir(APP_FILE_LOGGER)
	err := os.Mkdir(folderLog, os.ModePerm)
	if err != nil {
		Log = LogTemp
		LogTemp.Errorf("Error creating logger folder in %s, error is: %s\n", folderLog, err.Error())
		return
	}

	if _, err := os.OpenFile(APP_FILE_LOGGER, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm); err != nil {
		Log = LogTemp
		LogTemp.Errorf("Error creating logger file in %s, error is: %s\n", APP_FILE_LOGGER, err.Error())
		return
	}

	Log = logecs.NewLoggerEcs(logecs.EcsLogger{
		Mod: logAPi, Color: true,
		Path: APP_FILE_LOGGER, OutPut: true,
	})
}
