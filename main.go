package main

import (
	"new_whatsmeow/app"
	"new_whatsmeow/app/helpers/logg"

	"github.com/golang-module/carbon"
	"github.com/joho/godotenv"
)

func main() {
	logg.GeneralLogger.Printf("Starting app at %s ", carbon.Now().SubDays(40).Format("Y-m-d H:i:s"))

	err := godotenv.Load()
	if err != nil {
		logg.ErrorLogger.Println("\033[31mError loading env file\033[0m")
	}
	app.InitApp()
}
