package main

import (
	"go-crawler/common/mapp"
	"go-crawler/services/pkgGo"
)

// 注册服务
func SetUpServer() {
	pkgGo.Setup()
}

func main() {
	defer mapp.ServerClose()

	go SetUpServer()

	mapp.ServerRun()
}
