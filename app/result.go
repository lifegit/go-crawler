package app

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/lifegit/go-gulu/v2/nice/file"
	"github.com/lifegit/go-gulu/v2/pkg/ginMiddleware/mwCors"
	"github.com/lifegit/go-gulu/v2/pkg/ginMiddleware/mwLogger"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"strings"
	"time"
)

var Result ResultApi

type ResultApi struct {
	Gin *gin.Engine
	Api *gin.RouterGroup
}

func SetUpResult() {
	Result.Setup()
}

func (r *ResultApi) Setup() {
	// 设置模式，设置模式要放在调用Default()函数之前
	addr := fmt.Sprintf("%s:%d", Global.Server.Addr, Global.Server.Port)
	// mode
	gin.SetMode(Global.Server.RunMode)
	r.Gin = gin.New()
	// middleware
	// Recovery
	r.Gin.Use(gin.Recovery())
	// Cors
	if Global.Server.IsCors {
		r.Gin.Use(mwCors.NewCorsMiddleware())
	}
	// Logger
	r.Gin.Use(mwLogger.NewLoggerMiddlewareSmoothFail(true, Global.Server.HTTPLogDir))
	// staticPath
	if staticPath := Global.Server.StaticPath; staticPath != "" {
		Log.Info(fmt.Sprintf("visit http://%s/ for front-end static html files", addr))
		r.Gin.Use(static.Serve("/", static.LocalFile(staticPath, true)))
	}
	//if Global.Upload.Type == "local" {
	//	Log.Info(fmt.Sprintf("visit http://%s/%s for upload files", addr, path.Join(Global.Upload.Local.BaseDir)))
	//	r.Gin.StaticFS(Global.Upload.Local.BaseDir, http.Dir(Global.Upload.Local.BaseDir))
	//}
	// notRoute
	if Global.Server.IsNotRoute {
		Log.Info(fmt.Sprintf("visit http://%s/404 for RESTful Is NotRoute", addr))
		r.Gin.NoRoute(func(c *gin.Context) {
			c.Status(http.StatusNotFound)
		})
	}
	// swaggerApi
	if Global.Server.IsSwagger && Global.isDev() {
		Log.Info(fmt.Sprintf("visit http://%s/swagger/index.html for RESTful APIs Document", addr))
		// r.Gin.Static("swagger", "docs/swagger")
		r.Gin.GET("swagger/*any", func(c *gin.Context) {
			if strings.Contains(c.Request.RequestURI, "doc.json") {
				v, _ := file.ReadFile("docs/swagger/v3/openapi.json")
				_, _ = c.Writer.WriteString(v)
				c.Abort()
			}
		}, ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	// appInfo
	if Global.isDev() {
		Log.Info(fmt.Sprintf("visit http://%s/app/info for app info only on not-prod mode", addr))
		r.Gin.GET("/app/info", func(c *gin.Context) {
			c.JSON(http.StatusOK, Global)
		})
	}
	// apiPrefix
	r.Api = r.Gin.Group(Global.Server.APIPrefix)
}

func (r *ResultApi) Running() {
	addr := fmt.Sprintf("%s:%d", Global.Server.Addr, Global.Server.Port)
	Log.Infof("http result server at listening %s", addr)
	if Global.Server.IsHTTPS {
		// https
		if err := autotls.Run(r.Gin, addr); err != nil {
			Log.WithError(err).Fatal("https result server fail run !")
		}
		//if err := server.ListenAndServeTLS("cert.pem", "key.pem"); err != nil {
		//	Log.Errorf("https result server is run err: %v!", err)
		//}
	} else {
		// http
		server := &http.Server{
			Addr:           fmt.Sprintf(":%d", Global.Server.Port),
			Handler:        r.Gin,
			ReadTimeout:    time.Second * time.Duration(Global.Server.ReadTimeout),
			WriteTimeout:   time.Second * time.Duration(Global.Server.WriteTimeout),
			MaxHeaderBytes: 1 << 20,
		}
		if err := server.ListenAndServe(); err != nil {
			Log.Errorf("http result server is run err: %v!", err)
		}
	}

	//// endless
	//// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	//endless.DefaultReadTimeOut = readTimeout
	//endless.DefaultWriteTimeOut = writeTimeout
	//endless.DefaultMaxHeaderBytes = maxHeaderBytes
	//serverNew := endless.NewServer(endPoint, routersInit)
	//serverNew.BeforeBegin = func(add string) {
	//	Log.Info("Actual pid is %d", syscall.Getpid())
	//}
	//err = serverNew.ListenAndServe()
	//if err != nil {
	//	Log.Error("server err: %v", err)
	//}
}
