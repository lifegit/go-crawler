/**
* @Author: TheLife
* @Date: 2021/5/10 上午10:21
 */
package tasks

import (
	"fmt"
	"go-crawler/common/koa"
	"go-crawler/common/utils"
	"go-crawler/services/pkgGo/constant"
)

var printingTask *koa.Context

func init() {
	printingTask = utils.NewAweMediem(constant.ServiceName, "printing", printing)
}

var count = 0

func printing(k *koa.Context) {
	count++
	k.Result.Data = count
	fmt.Println("task examples")
}
