/**
* @Author: TheLife
* @Date: 2021/5/10 上午10:20
 */
package tasks

import (
	"fmt"
	"go-crawler/common/appLogging"
	"go-crawler/common/mediem"
	"go-crawler/common/mediem/midMiddleware"
	"go-crawler/services/pkgGo/constant"
	"go-gulu/core"
)

var tasks *core.Scheduler
var loggerMid mediem.HandlerFunc

func init() {
	tasks = core.NewScheduler()

	loggerM, err := midMiddleware.NewLoggerMiddleware(true, true, constant.AppName, fmt.Sprintf("./runtime/logs/services/%s/examples", constant.AppName))
	if err != nil {
		appLogging.Log.WithError(err).Fatal(fmt.Sprintf("services server fail run %s !", constant.AppName), err)
	}
	loggerMid = loggerM
}

func RunTasks() {
	tasks.Every(5).Seconds().Do(printing).Run()
	tasks.Every(5).Seconds().Do(search)

	<-tasks.Start()
}
