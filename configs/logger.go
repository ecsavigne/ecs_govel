package configs

import (
	"fmt"
	"log"
	"os"
)

var (
	GeneralLogger *log.Logger
	// ErrorLogger exported
	ErrorLogger *log.Logger
	// ErrorLogger exported
	WarningLogger *log.Logger
	// AlertLogger exported
	AlertLogger *log.Logger
	// FatalLogger exported
	FatalLogger *log.Logger
)

func prepare_logger() {
	generalLog, err := os.OpenFile(APP_FILE_LOGGER, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	GeneralLogger = log.New(generalLog, "General Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)

	ErrorLogger = log.New(generalLog, "Error Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)

	WarningLogger = log.New(generalLog, "Warning Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)

	AlertLogger = log.New(generalLog, "Alert Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)

	FatalLogger = log.New(generalLog, "Fatal Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
}
