package schedular

import (
	"reviewservice/config/env"

	"github.com/robfig/cron/v3"
)


func SetupCron(){
	c := cron.New() 

	c.AddFunc(env.GetString("CRON_EXPR" , "* * * * *") , func() {

	})
}