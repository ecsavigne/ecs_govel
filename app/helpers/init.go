package helpers

import (
	"new_whatsmeow/app/services/gdrive"

	"github.com/joho/godotenv"
)

var srv = new(gdrive.ServiceGoogle)

func init() {
	godotenv.Load()
}
