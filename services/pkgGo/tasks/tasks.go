/**
* @Author: TheLife
* @Date: 2021/5/10 上午10:20
 */
package tasks

import (
	"go-gulu/core"
)

func Run() {
	tasks := core.NewScheduler()
	tasks.Every(5).Seconds().Do(func() { printingTask.Run() }).Run()
	tasks.Every(8).Seconds().Do(func() { searchTask.Run() })

	<-tasks.Start()
}
