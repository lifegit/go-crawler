package pkgGo

import (
	"fmt"
	"go-crawler/app"
	"go-crawler/services/pkgGo/constant"
	"go-crawler/services/pkgGo/handlers/v1"
	"go-crawler/services/pkgGo/tasks"
)

func Setup() {
	app.Log.Info(fmt.Sprintf("service run %s", constant.ServiceName))

	v1.Run()
	go tasks.Run()
}
