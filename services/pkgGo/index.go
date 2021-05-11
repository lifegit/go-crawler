/**
* @Author: TheLife
* @Date: 2021/5/8 下午5:58
 */
package pkgGo

import (
	"fmt"
	"go-crawler/common/appLogging"
	"go-crawler/services/pkgGo/constant"
	_ "go-crawler/services/pkgGo/handlers/v1"
	"go-crawler/services/pkgGo/tasks"
)

func Setup() {
	appLogging.Log.Info(fmt.Sprintf("service run %s", constant.AppName))

	go tasks.RunTasks()
}
