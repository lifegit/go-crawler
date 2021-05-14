/**
* @Author: TheLife
* @Date: 2021/5/10 上午10:05
 */
package v1

import (
	"go-crawler/common/utils"
	"go-crawler/services/pkgGo/constant"
)

func Run() {
	v1 := utils.NewAweHandlers(constant.ServiceName, "v1")
	{
		v1.GET("example", example)
	}
}
