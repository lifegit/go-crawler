/**
* @Author: TheLife
* @Date: 2021/5/13 上午10:18
 */
package utils

import (
	"github.com/gin-gonic/gin"
	"go-crawler/common/appLogging"
	"go-crawler/common/conf"
	"go-crawler/common/mapp"
	"go-crawler/common/mediem"
	"go-crawler/common/mediem/midMiddleware"
	"go-gulu/ginMiddleware/mwLogger"
	"path"
)

func NewAweMediem(serviceName, methodName string, workHandlerFunc mediem.HandlerFunc) *mediem.Context {
	loggerMid := midMiddleware.NewLoggerMiddlewareSmoothFail(true, true, serviceName, path.Join(conf.GetString("service.log"), serviceName, "task", methodName), appLogging.Log)
	printingTask := mediem.NewContext()
	printingTask.Use(midMiddleware.Recovery(), workHandlerFunc, loggerMid)

	return printingTask
}

func NewAweHandlers(serviceName, version string) *gin.RouterGroup {
	loggerMid := mwLogger.NewLoggerMiddlewareSmoothFail(true, true, path.Join(conf.GetString("service.log"), serviceName, "http"), appLogging.Log)

	return mapp.Result.Api.Group(serviceName).Group(version, loggerMid)
}
