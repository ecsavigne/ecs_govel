package app

// import (
// 	"context"
// 	"flag"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"path/filepath"
// 	"strings"
// 	"time"

// 	"ecs_govel/app/helpers/logg"
// 	"ecs_govel/config"

// 	configDataBase "ecs_govel/config"
// 	"ecs_govel/routes"

// 	"github.com/go-co-op/gocron"
// 	"github.com/joho/godotenv"
// )

// var AppConfig *configDataBase.Config

func InitApp() {
	// // 1. handle exceptions
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		//	logg.Log("[App]", "", "", "Recovered from exception. Interface in defer is: "+fmt.Sprintf("%+v", err), true)
	// 	}
	// }()
	// godotenv.Load()

	// // 2. safe shutdown
	// var wait time.Duration
	// flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	// flag.Parse()

	// // 3. init database
	// if os.Getenv("ecs_govel/app_ENV") == "local" {
	// 	/*seeders.Seeder.ApplicationSeeder()
	// 	seeders.Seeder.CompaniesWhatsappSeeder()*/
	// }

	// // 5. load config.yml
	// AppConfig = configDataBase.NewConfig(config.ConfigPath)

	// // 6. update port if executavel have the format: main[port_number]
	// filename := filepath.Base(os.Args[0])
	// port := strings.Replace(filename, "main", "", 1)
	// if port == "__debug_bin" || port == "" { // For debugg purposse
	// 	port = "1337"
	// 	AppConfig.Server.AppID = "test"
	// } else if port != "" {
	// 	AppConfig.Server.Port = port
	// 	AppConfig.Server.AppID = strings.Replace(AppConfig.Server.Port, "1337", "", 1)
	// 	AppConfig.Server.AppID = strings.Replace(AppConfig.Server.AppID, "133", "", 1)
	// }
	// configDataBase.Database.AppID = AppConfig.Server.AppID

	// // 7. create a gocron scheduler
	// scheduler := gocron.NewScheduler(time.UTC)
	// // if port == "13371" {
	// // 	_, _ = scheduler.Every(120).Seconds().Do(controllers.MessengerController.CallRouter)
	// // 	scheduler.StartAsync()
	// // }

	// // 8. Restore sessions of already loggued numbers
	// //repositories.SessionRepository.RestoreAssignedClients()

	// // 9. http server configuration for this main process
	// server := &http.Server{
	// 	Handler:      routes.Router,
	// 	Addr:         AppConfig.Server.Host + ":" + AppConfig.Server.Port,
	// 	WriteTimeout: time.Duration(120) * time.Second,
	// 	ReadTimeout:  time.Duration(AppConfig.Server.Timeout.Read) * time.Second,
	// 	IdleTimeout:  time.Duration(AppConfig.Server.Timeout.Idle) * time.Second,
	// }

	// // 10. run the http server paralelly in a goroutine to receive request
	// go func() {
	// 	if err := server.ListenAndServe(); err != nil {
	// 		logMessage := "Impossible ListenAndServe: " + err.Error()
	// 		logg.Log("[App]", "", "", logMessage, true)
	// 	}
	// }()

	// // 11. print a logg message
	// //logMessage := "Server Addr: " + AppConfig.Server.Host + ":" + AppConfig.Server.Port
	// //logg.Log("[App]", "", "", logMessage, true)

	// // 12. hanle Ctrl+C shuthdown of this main process
	// c := make(chan os.Signal, 1)   //accept graceful shutdowns when quit via SIGINT
	// signal.Notify(c, os.Interrupt) // Block until we receive our signal.
	// <-c
	// ctx, cancel := context.WithTimeout(context.Background(), wait) // Create a deadline to wait for.
	// defer cancel()
	// _ = server.Shutdown(ctx) // Doesn't block if no connections, but will otherwise wait until the timeout deadline.
	// logg.Log("[Main]", "", "", "shutting down", true)
	// scheduler.Clear()
	// os.Exit(0)
}
