package tasks

import (
	"fmt"
	"github.com/lifegit/go-gulu/v2/nice/koa"
	"go-crawler/app"
	"go-crawler/services/pkgGo/constant"
)

var printingTask *koa.Context

func init() {
	printingTask = app.NewAweMediem(constant.ServiceName, "printing", printing)
}

var count = 0

func printing(k *koa.Context) {
	count++
	k.Result.Data = count
	fmt.Println("task examples")
}
