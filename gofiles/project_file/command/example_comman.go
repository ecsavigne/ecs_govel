package command

import "github.com/spf13/viper"

func (self *command) CmdExemple() {
	interval := viper.GetString("TEST_JOB_INTERVAL")
	dateBegin := viper.GetString("TEST_JOB_DATE")
	key := "TEST_JOB"

	self.createCmd(key, interval, dateBegin, func() {
		logs.Sub("Command").Warnf("Executing method: 'CmdExemple'\n")
	})
}
