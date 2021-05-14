/**
* @Author: TheLife
* @Date: 2020-10-30 2:37 上午
 */
package mapp

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"go-crawler/common/appLogging"
	"go-crawler/common/conf"
	"go-gulu/ginMiddleware/mwCors"
	"net/http"
	"path"
	"time"
)

var Result ResultApi

type ResultApi struct {
	Gin *gin.Engine
	Api *gin.RouterGroup
}

func (r *ResultApi) Setup() {
	// 设置模式，设置模式要放在调用Default()函数之前
	addr := getAddr()

	// mode
	gin.SetMode(conf.GetString("server.runMode"))
	r.Gin = gin.New()
	// middlewareRecovery
	r.Gin.Use(gin.Recovery())
	// middlewareCors
	if conf.GetBool("server.isCors") {
		r.Gin.Use(mwCors.NewCorsMiddleware())
	}
	// staticPath
	if staticPath := conf.GetString("server.staticPath"); staticPath != "" {
		appLogging.Log.Info(fmt.Sprintf("visit http://%s/ for front-end static html files", addr))
		r.Gin.Use(static.Serve("/", static.LocalFile(staticPath, true)))
	}
	// notRoute
	if conf.GetBool("server.isNotRoute") {
		appLogging.Log.Info(fmt.Sprintf("visit http://%s/docs for RESTful Is NotRoute", addr))
		r.Gin.NoRoute(func(c *gin.Context) {
			//file := path.Join(sp, "index.html")
			//c.File(file)
			c.Status(http.StatusNotFound)
		})
	}
	// swaggerApi
	if conf.GetBool("server.isSwagger") && conf.GetString("app.env") != "prod" {
		appLogging.Log.Info(fmt.Sprintf("visit http://%s/docs for RESTful APIs Document", addr))
		//add edit your own swagger.docs.yml file in ./swagger/docs.yml
		//generateSwaggerDocJson()
		r.Gin.Static("docs", "./docs")
	}
	// appInfo
	if conf.GetString("app.env") != "prod" {
		appLogging.Log.Info(fmt.Sprintf("visit http://%s/app/info for app info only on not-prod mode", addr))
		r.Gin.GET("/app/info", func(c *gin.Context) {
			m := make(map[string]map[string]string)
			for _, val := range []string{"app", "server", "enable", "upload"} {
				m[val] = conf.GetStringMapString(val)
			}
			c.JSON(200, m)
		})
	}

	// apiPrefix
	r.Api = r.Gin.Group(path.Join("services", conf.GetString("server.apiPrefix")))
}

func (r *ResultApi) Running() {
	addr := getAddr()
	appLogging.Log.Infof("http result server at listening %s", addr)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.GetInt("server.port")),
		Handler:        r.Gin,
		ReadTimeout:    time.Second * time.Duration(conf.GetInt("server.readTimeout")),
		WriteTimeout:   time.Second * time.Duration(conf.GetInt("server.writeTimeout")),
		MaxHeaderBytes: 1 << 20,
	}
	if conf.GetBool("server.isHttps") {
		// https
		if err := autotls.Run(r.Gin, addr); err != nil {
			appLogging.Log.WithError(err).Fatal("https result server fail run !")
		}
		//if err := server.ListenAndServeTLS("cert.pem", "key.pem"); err != nil {
		//	appLogging.Log.Errorf("https result server is run err: %v!", err)
		//}
	} else {
		// http
		if err := server.ListenAndServe(); err != nil {
			appLogging.Log.Errorf("http result server is run err: %v!", err)
		}
	}

	//// endless
	//// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	//endless.DefaultReadTimeOut = readTimeout
	//endless.DefaultWriteTimeOut = writeTimeout
	//endless.DefaultMaxHeaderBytes = maxHeaderBytes
	//serverNew := endless.NewServer(endPoint, routersInit)
	//serverNew.BeforeBegin = func(add string) {
	//	logging.Logger.Info("Actual pid is %d", syscall.Getpid())
	//}
	//err = serverNew.ListenAndServe()
	//if err != nil {
	//	appLogging.Logger.Error("server err: %v", err)
	//}
}

func getAddr() string {
	return fmt.Sprintf("%s:%d", conf.GetString("server.addr"), conf.GetInt("server.port"))
}
