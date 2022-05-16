package main

import (
	"go-crawler/app"
	"go-crawler/services/pkgGo"
)

// SetUpServer 注册服务
func SetUpServer() {
	pkgGo.Setup()
}

func main() {
	defer app.Close()

	app.Log.Infof("app: %s , version: %d at start ...", app.Global.App.Name, app.Global.App.Version)
	go SetUpServer()

	app.Run()
}
