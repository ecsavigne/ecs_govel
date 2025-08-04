package configs

import (
	"fmt"
	"log"
	"os"
	"path"

	logecs "github.com/ecsavigne/logecs/log"
)

// log
var (
	Log, LogTemp logecs.Logger
	f            *os.File
)

func prepare_logger() {
	dir := path.Dir(APP_FILE_LOGGER)
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf(`
			Panic: %+v

			%sPlease create folder manually and run the application again.
			1. sudo mkdir -p %s%s%s%s. If exists, please run step 2 and run the application again.
			2. sudo chmod -R 777 %s %s.
			`, r, "\x1b[34m", "\x1b[0m", "\x1b[31m", dir, "\x1b[0m\x1b[34m", dir, "\x1b[0m")
		}
	}()

	// create folder if not exists
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating logger file in %s, error is: %s\n", APP_FILE_LOGGER, err.Error())
	}

	// open file if not exists and append data or create
	f, err = os.OpenFile(APP_FILE_LOGGER, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		Log.Errorf("Error creating logger file in %s, error is: %s\n", APP_FILE_LOGGER, err.Error())
	}

	Log = logecs.NewLoggerEcs(logecs.EcsLogger{
		Mod: logAPi, Color: true,
		Path: APP_FILE_LOGGER, OutPut: true,
	})

	Log.Infof("\x1b[34mConfiguring logger file in %s\x1b[0m\n", APP_FILE_LOGGER)
}
