package tasks

import "github.com/lifegit/go-gulu/v2/nice/core"

func Run() {
	tasks := core.NewScheduler()
	tasks.Every(5).Seconds().Do(func() { printingTask.Run() }).Run()
	tasks.Every(8).Seconds().Do(func() { searchTask.Run() })

	<-tasks.Start()
}
