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
	"fmt"
	"go-crawler/common/appLogging"
	"go-crawler/{{.ServiceDir}}/{{.ServiceName}}/constant"
	"go-crawler/{{.ServiceDir}}/{{.ServiceName}}/handlers/v1"
	"go-crawler/{{.ServiceDir}}/{{.ServiceName}}/tasks"
)

func Setup() {
	appLogging.Log.Info(fmt.Sprintf("service run %s", constant.ServiceName))

	v1.Run()
	go tasks.Run()
}
`,
}

var tConstantIndex = tplNode{
	"constant/constant.go",
	`
package constant

const ServiceName = "{{.ServiceName}}"
`,
}

var tHandlersAppIndex = tplNode{
	"handlers/v1/v1.go",
	`
package v1

import (
	"go-crawler/common/utils"
	"go-crawler/{{.ServiceDir}}/{{.ServiceName}}/constant"
)

func Run()  {
	v1 := utils.NewAweHandlers(constant.ServiceName, "v1")
	{
		v1.GET("example", example)
	}
}
`,
}

var tHandlersExampleIndex = tplNode{
	"handlers/v1/example.go",
	`
package v1

import (
	"github.com/gin-gonic/gin"
	"go-gulu/app"
)

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
	"fmt"
	"go-crawler/common/mediem"
	"go-crawler/common/utils"
	"go-crawler/{{.ServiceDir}}/{{.ServiceName}}/constant"
)

var printingTask *mediem.Context

func init() {
	printingTask = utils.NewAweMediem(constant.ServiceName, "printing", printing)
}


var count = 0
func printing(c *mediem.Context) {
	count++
	c.Result.Data = count
	fmt.Println("task examples")
}
`,
}

var tTasksSearchIndex = tplNode{
	"tasks/search.go",
	`
/**
* @Author: TheLife
* @Date: 2021/5/10 上午10:21
 */
package tasks

import (
	"github.com/chromedp/chromedp"
	"go-crawler/common/mediem"
	"go-crawler/common/utils"
	"go-crawler/{{.ServiceDir}}/{{.ServiceName}}/constant"
)

var searchTask *mediem.Context

func init() {
	searchTask = utils.NewAweMediem(constant.ServiceName, "search", search)
}


func search(c *mediem.Context) {
	// 1. create chrome instance
	ctx, cancel := utils.NewAweChromeDp(10, true)
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

	c.Result.Data = example
	c.Result.Err = err
}
`,
}

var tTasksTaskIndex = tplNode{
	"tasks/task.go",
	`
package tasks

import (
	"go-gulu/core"
)

func Run() {
	tasks := core.NewScheduler()
	tasks.Every(5).Seconds().Do(func () {printingTask.Run()}).Run()
	tasks.Every(8).Seconds().Do(func() {searchTask.Run()})

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
