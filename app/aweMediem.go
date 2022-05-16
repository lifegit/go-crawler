package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lifegit/go-gulu/v2/nice/koa"
	"github.com/lifegit/go-gulu/v2/nice/koa/koaMiddleware"
	"github.com/lifegit/go-gulu/v2/pkg/ginMiddleware/mwLogger"
	"path"
)

func NewAweMediem(serviceName, methodName string, workHandlerFunc koa.HandlerFunc) *koa.Context {
	return koa.NewContext().Use(
		koaMiddleware.Recovery(),
		workHandlerFunc,
		koaMiddleware.NewLoggerMiddlewareSmoothFail(true, true, serviceName, path.Join(Global.Server.HTTPLogDir, serviceName, "task", methodName)),
	)
}

func NewAweHandlers(serviceName, version string) *gin.RouterGroup {
	return Result.Api.
		Group(serviceName).
		Group(version, mwLogger.NewLoggerMiddlewareSmoothFail(true, path.Join(Global.Server.HTTPLogDir, serviceName, "http")))
}
