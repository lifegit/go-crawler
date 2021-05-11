/**
* @Author: TheLife
* @Date: 2021/5/10 上午10:05
 */
package v1

import (
	"github.com/gin-gonic/gin"
	"go-crawler/common/mapp"
	"go-crawler/services/pkgGo/constant"
)

var api *gin.RouterGroup

func init() {
	api = mapp.Result.Api.Group(constant.AppName).Group("v1")
}
