# 微爬虫框架

### 介绍

 一个项目中挂载多个微小服务（互相相不影响），各项服务可分别获取多个友商的开放数据，从而建立行业内大数据平台。


### 项目结构

    crawler
    |-- common                                    # 一些常用工具
    |   |-- appLogging                            # 整个应用的log
    |   |-- mapp                                  # 整个应用的工具集合，包含初始化、路由、等等 
    |   |-- mediem                                # koa 洋葱中间件模型，思想同gin的Context。可用于🕷️的中间件。
    |   |-- spider                                # 爬虫的基础工具，chromedp 、colly 等
    |   |-- utils                                 # 全局工具类
    |-- conf                                      # 配置文件
    |-- docs                                      # 文档  
    |-- ipManager                                 # Ip 池         
    |-- scripts                                   # 脚本 
    |-- services                                  # 服务   
        |-- serviceName                           # 某项微服务
             |-- index                            # 服入口，通常用于初始化服务
             |-- constant                         # 该服务的一些常量
             |-- handlers                         # 该服务对外的开放Api（http、rpc）
             |-- models                           # 持久化模型
             |-- tasks                            # 🕷️数据源获取
                          
### 使用

    假设服务名为 pkgGo，以该服务为介绍：
    
#### 初始化
    
    1. 可使用 scripts/newServices 来生成一个服务，或自己实现服务。
    2. 在 main.go 的 SetUpServer 函数中注册服务。
     ```
        func SetUpServer()  {
        	pkgGo.Setup()
        }
     ```
    
#### 服务
    1. 在该服务的入口文件：`services/pkgGo/index.go` 可初始化 对外Api 、 🕷️ 服务 等。
    2. 🕷️ 服务推荐使用 koa 包的洋葱中间件模型，可共用一些内置中间件来解决额外的问题。
    ```
    	var k koa.Context
    	k.Use(koaMiddleware.Recovery(), fetch, loggerMid).Run()
    3. ip池(拿到可用的代理Ip)：ipManager、持久化：mongodb、缓存：redis，等 可自行根据业务使用。
    ```

### 工具
> ⚠️ 暂未完成
- 缓存：redis
- 持久化：mongodb ⚠️
- ip池：ipManager ⚠️
- 数据清洗：ipManager ⚠️
- 单元测试：⚠️


### 感谢
- [gin-gonic/gin](https://github.com/gin-gonic/gin)
- [chromedp](https://github.com/chromedp/chromedp)
- [colly](https://github.com/gocolly/colly)
- [logrus](https://github.com/sirupsen/logrus)
- [viper](https://github.com/spf13/viper)
- [go-redis](https://github.com/go-redis/redis)
- [qmgo](https://github.com/qiniu/qmgo)