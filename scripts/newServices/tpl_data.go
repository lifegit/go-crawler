/**
* @Author: TheLife
* @Date: 2021/5/11 上午9:05
 */
package newServices

var tIndex = tplNode{
	"index.go",
	`
package {{.ServiceName}}

import (
	"go-crawler/common/appLogging"
	"go-crawler/{{.ServiceDir}}/{{.ServiceName}}/constant"
	_ "go-crawler/{{.ServiceDir}}/{{.ServiceName}}/handlers/v1"
	"go-crawler/{{.ServiceDir}}/{{.ServiceName}}/tasks"
	"fmt"
)

func Setup() {
	appLogging.Log.Info(fmt.Sprintf("service run %s", constant.AppName))

	go tasks.RunTasks()
}
`,
}

var tConstantIndex = tplNode{
	"constant/constant.go",
	`
package constant


const AppName = "{{.ServiceName}}"
`,
}

var tHandlersAppIndex = tplNode{
	"handlers/v1/h_app.go",
	`
package v1

import (
	"go-crawler/common/mapp"
	"go-crawler/{{.ServiceDir}}/{{.ServiceName}}/constant"
	"github.com/gin-gonic/gin"
)

var api *gin.RouterGroup

func init() {
	api = mapp.Result.Api.Group(constant.AppName).Group("v1")
}
`,
}

var tHandlersExampleIndex = tplNode{
	"handlers/v1/h_example.go",
	`
package v1

import (
	"github.com/gin-gonic/gin"
	"go-gulu/app"
)


func init() {
	api.GET("example", example)
}

//All
func example(c *gin.Context) {
	app.JsonSuccess(c)
}
`,
}

var tTasksPrintingIndex = tplNode{
	"tasks/printing.go",
	`
package tasks

import (
	"go-crawler/common/mediem"
	"go-crawler/common/mediem/midMiddleware"
	"fmt"
)

var count = 0
func printing()  {
	fetch := func(c *mediem.Context) {
		count++
		c.Result.Data = count
		fmt.Println("task examples")
	}

	var me mediem.Context
	me.Use(midMiddleware.Recovery(), fetch, loggerMid).Run()
}
`,
}

var tTasksSearchIndex = tplNode{
	"tasks/search.go",
	`
package tasks

import (
	"go-crawler/common/mediem"
	"go-crawler/common/mediem/midMiddleware"
	"go-crawler/common/spider/chromedpp"
	"github.com/chromedp/chromedp"
)


func search()  {
	fetch := func(c *mediem.Context) {
		// 1. create chrome instance
		ctx, cancel := chromedpp.NewChromeDp(10, true)
		defer cancel()

		// 2. search baidu
		var example string
		err := chromedp.Run(*ctx,
			chromedp.Navigate("https://pkg.go.dev/time"),
	// wait for footer element is visible (ie, page is loaded)
	chromedp.WaitVisible("body > footer"),
	// find and click "Example" link
	chromedp.Click("#example-After", chromedp.NodeVisible),
	// retrieve the text of the textarea
	chromedp.Value("#example-After textarea", &example),
)

	c.Result.Err = err
}

var me mediem.Context
me.Use(midMiddleware.Recovery(), fetch, loggerMid).Run()
}
`,
}

var tTasksTaskIndex = tplNode{
	"tasks/task.go",
	`
package tasks

import (
	"go-crawler/common/appLogging"
	"go-crawler/common/mediem"
	"go-crawler/common/mediem/midMiddleware"
	"go-crawler/{{.ServiceDir}}/{{.ServiceName}}/constant"
	"fmt"
	"go-gulu/core"
)

var tasks *core.Scheduler
var loggerMid mediem.HandlerFunc

func init() {
	tasks = core.NewScheduler()

	loggerM, err := midMiddleware.NewLoggerMiddleware(true, true, constant.AppName, fmt.Sprintf("./runtime/logs/{{.ServiceDir}}/%s/examples", constant.AppName))
	if err != nil{
		appLogging.Log.WithError(err).Fatal(fmt.Sprintf("services server fail run %s !", constant.AppName), err)
	}
	loggerMid = loggerM
}


func RunTasks()  {
	tasks.Every(5).Seconds().Do(printing).Run()
	tasks.Every(5).Seconds().Do(search)

	<-tasks.Start()
}
`,
}

var parseOneList = []tplNode{
	tIndex,
	tConstantIndex,
	tHandlersAppIndex, tHandlersExampleIndex,
	tTasksPrintingIndex, tTasksSearchIndex, tTasksTaskIndex,
}
