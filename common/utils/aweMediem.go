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
	"go-crawler/common/koa"
	"go-crawler/common/koa/koaMiddleware"
	"go-gulu/ginMiddleware/mwLogger"
	"path"
)

func NewAweMediem(serviceName, methodName string, workHandlerFunc koa.HandlerFunc) *koa.Context {
	loggerMid := koaMiddleware.NewLoggerMiddlewareSmoothFail(true, true, serviceName, path.Join(conf.GetString("service.log"), serviceName, "task", methodName), appLogging.Log)
	printingTask := koa.NewContext()
	printingTask.Use(koaMiddleware.Recovery(), workHandlerFunc, loggerMid)

	return printingTask
}

func NewAweHandlers(serviceName, version string) *gin.RouterGroup {
	loggerMid := mwLogger.NewLoggerMiddlewareSmoothFail(true, true, path.Join(conf.GetString("service.log"), serviceName, "http"), appLogging.Log)

	return mapp.Result.Api.Group(serviceName).Group(version, loggerMid)
}
