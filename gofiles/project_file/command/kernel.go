package command

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"time"

	logecs "github.com/ecsavigne/logecs/log"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
)

type timesProgramsCmd struct {
	interval   string    // ex: 1h1m1s
	date       time.Time // ex: 2023-01-01T00:00:00Z, date for job begin, if date is not set, the job will be executed immediately
	is_execute bool      // if true, the job be executed
	idJobs     cron.EntryID
}

var (
	logs  logecs.Logger
	cron_ = cron.New()
	// ids is a map that stores the id of the job in the cron, it is used to remove the job when the time is changed
	// @key = nome of id in ids map representing the job located in .env session "Time per job for app shedule", @value = id of the job in the cron
	ids map[string]cron.EntryID = make(map[string]cron.EntryID)
	// timeJob is the time for the job to be executed, it is set in the config_env.env file
	// @key = nome of id in ids map representing the job located in .env session "Time per job for app shedule", @value = time for the job to be executed
	timesJobs = make(map[string]*timesProgramsCmd)
)

type command struct{}

func SetLogger(l logecs.Logger) {
	logs = l
}

func isProcessPrincipal(processNumber string) bool {
	if processNumber == viper.GetString("APP_NAME_X1") {
		return true
	}
	return false
}

func dateValidForInitial(date time.Time) bool {
	yP, mP, dP := date.Date()
	hP, minP, _ := date.Clock()

	now := time.Now()
	yNow, mNow, dNow := now.Date()
	hNow, minNow, _ := now.Clock()

	if yP == yNow && mP == mNow && dP == dNow && hP == hNow && minP == minNow {
		return true
	}

	return false
}

//	create one job
//
// @keyJob       = key of the job in the config file
// @interval     = time for the job to be executed
// @dateBeginCmd = date for the job to be executed
// @cmd 		  = function to be executed
func (*command) createCmd(keyJob, interval, dateBeginCmd string, cmd func()) {
	dateBegin, _ := time.Parse(time.DateTime, dateBeginCmd)

	if infoJobs, ok := timesJobs[keyJob]; ok {
		if infoJobs.is_execute && infoJobs.idJobs != 0 {
			if infoJobs.interval != interval {
				logs.Sub("Command").Infof("Removing job: %s with id: %d\n", keyJob, infoJobs.idJobs)
				cron_.Remove(infoJobs.idJobs)

				infoJobs.interval = interval
				id, e := cron_.AddFunc(fmt.Sprintf("@every %s", infoJobs.interval), cmd)
				infoJobs.idJobs = id
				if e != nil {
					logs.Sub("Command").Errorf("Error adding job: %s, id_job: %d, with interval: %s, error: %v\n", keyJob, infoJobs.idJobs, interval, e)
					return
				}
			}
		}
	} else {
		if dateValidForInitial(dateBegin) || dateBeginCmd == "" {
			id, e := cron_.AddFunc(fmt.Sprintf("@every %s", interval), cmd)
			if e != nil {
				logs.Sub("Command").Errorf("Error adding job: %s with time: %s, error: %v\n", keyJob, interval, e)
				return
			}
			timesJobs[keyJob] = &timesProgramsCmd{interval: interval, date: dateBegin, idJobs: id, is_execute: true}
			logs.Sub("Command").Infof("Create job : %s with time: %s\n", keyJob, interval)
			return
		}
	}
}

func shedulerLive() {
	commands := command{}

	c := reflect.TypeFor[*command]()
	value := reflect.ValueOf(&commands)

	for method := range c.Methods() {
		value.MethodByName(method.Name).Call([]reflect.Value{})
	}
}

func RunJobs() {
	defer func() {
		if r := recover(); r != nil {
			logMessage := "Error in;.........  " + fmt.Sprintf("%+v", r)
			// panic(logMessage)
			logs.Sub("Command").Errorf("[ RunJobs ] - %s", logMessage)
		}
	}()

	logs.Sub("Command").Debugf("Starting cron jobs... : %s \n", viper.GetString("APP_NAME_X1"))
	processName := filepath.Base(os.Args[0])
	if !isProcessPrincipal(processName) {
		return
	}

	cron_.AddFunc("@every 0h0m3s", shedulerLive)

	cron_.Start()
}
