/**
* @Author: TheLife
* @Date: 2021/5/10 上午10:21
 */
package tasks

import (
	"fmt"
	"go-crawler/common/mediem"
	"go-crawler/common/mediem/midMiddleware"
)

var count = 0

func printing() {
	fetch := func(c *mediem.Context) {
		count++
		c.Result.Data = count
		fmt.Println("task examples")
	}

	var me mediem.Context
	me.Use(midMiddleware.Recovery(), fetch, loggerMid).Run()
}
