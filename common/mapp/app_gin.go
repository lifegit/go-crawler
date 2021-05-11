/**
* @Author: TheLife
* @Date: 2020-10-30 2:37 上午
 */
package mapp

import (
	"errors"
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"go-crawler/common/appLogging"
	"go-crawler/common/conf"
	"go-gulu/ginMiddleware/mwCors"
	"go-gulu/ginMiddleware/mwJwt"
	"go-gulu/ginMiddleware/mwLogger"
	"net/http"
	"path"
	"time"
)

var Result ResultApi

type ResultApi struct {
	Gin *gin.Engine
	Api *gin.RouterGroup

	Jwt  mwJwt.MwJwt
	Addr string
}

func (r *ResultApi) Setup() {
	// 设置模式，设置模式要放在调用Default()函数之前
	r.Addr = fmt.Sprintf("%s:%d", conf.GetString("server.addr"), conf.GetInt("server.port"))

	// mode
	gin.SetMode(conf.GetString("server.runMode"))
	r.Gin = gin.New()
	// middlewareRecovery
	r.Gin.Use(gin.Recovery())
	// middlewareLogger
	middlewareLogger, err := mwLogger.NewLoggerMiddleware(true, conf.GetBool("enable.httpLog"), conf.GetString("server.logDir"))
	if err != nil {
		appLogging.Log.WithError(err).Fatal("gin middleware logger is io error")
	}
	r.Gin.Use(middlewareLogger)
	// middlewareCors
	if conf.GetBool("enable.cors") {
		r.Gin.Use(mwCors.NewCorsMiddleware())
	}
	// middlewareJwt
	r.Jwt = mwJwt.NewJwtMiddleware(conf.GetString("jwt.key"), conf.GetString("app.name"), conf.GetString("jwt.secret"), conf.GetString("jwt.key"))

	// staticPath
	if staticPath := conf.GetString("server.staticPath"); staticPath != "" {
		appLogging.Log.Info(fmt.Sprintf("visit http://%s/ for front-end static html files", r.Addr))
		r.Gin.Use(static.Serve("/", static.LocalFile(staticPath, true)))
	}
	// notRoute
	if conf.GetBool("enable.notRoute") {
		appLogging.Log.Info(fmt.Sprintf("visit http://%s/docs for RESTful Is NotRoute", r.Addr))
		r.Gin.NoRoute(func(c *gin.Context) {
			//file := path.Join(sp, "index.html")
			//c.File(file)
			c.Status(http.StatusNotFound)
		})
	}
	// swaggerApi
	if conf.GetBool("enable.swagger") && conf.GetString("app.env") != "prod" {
		appLogging.Log.Info(fmt.Sprintf("visit http://%s/docs for RESTful APIs Document", r.Addr))
		//add edit your own swagger.docs.yml file in ./swagger/docs.yml
		//generateSwaggerDocJson()
		r.Gin.Static("docs", "./docs")
	}
	// appInfo
	if conf.GetString("app.env") != "prod" {
		appLogging.Log.Info(fmt.Sprintf("visit http://%s/app/info for app info only on not-prod mode", r.Addr))
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
	appLogging.Log.Infof("http result server at listening %s", r.Addr)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.GetInt("server.port")),
		Handler:        r.Gin,
		ReadTimeout:    time.Second * time.Duration(conf.GetInt("server.readTimeout")),
		WriteTimeout:   time.Second * time.Duration(conf.GetInt("server.writeTimeout")),
		MaxHeaderBytes: 1 << 20,
	}
	if conf.GetBool("enable.https") {
		// https
		if err := autotls.Run(r.Gin, r.Addr); err != nil {
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

type JwtUser struct {
	Id       uint
	Username string
}

func (r *ResultApi) GetJwtUser(c *gin.Context) (user *JwtUser, err error) {
	//return &utils.User{
	//	Id:1,
	//	Username:"12345678",
	//},nil

	res, exists := c.Get(conf.GetString("jwt.key"))
	if !exists || res == nil {
		return nil, errors.New("参数不存在")
	}

	if err := mapstructure.Decode(res, &user); err != nil {
		return nil, errors.New("参数不存在")
	}

	return
}
