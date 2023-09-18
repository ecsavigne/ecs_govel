package logg

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var GeneralLogger *log.Logger
var ErrorLogger *log.Logger

func init() {
	generalLog, err := os.OpenFile("./logs/serverWeb.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	GeneralLogger = log.New(generalLog, "General Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(generalLog, "Error Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
}

func ProcessError(err error, AppID string) {
	pos := strings.Index(err.Error(), "too many")

	if pos > -1 {
		GeneralLogger.Println("[Running Pkill by Too many Error]")

		setpassword := "edson2021"
		_, err := exec.Command("sh", "-c", "echo '"+setpassword+"' | sudo -S pkill -SIGINT 'main1337"+AppID+"'").Output()

		if err != nil {
			fmt.Println("[Pkill by Too many Error]: Falha ao finalizar processo")
			GeneralLogger.Println("[Pkill by Too many Error]: Falha ao finalizar processo")

			return
		}

		GeneralLogger.Println("[Apllication restarting...!!!]")
	}
}
