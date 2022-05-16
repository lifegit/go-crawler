module go-crawler

go 1.15

require (
	github.com/chromedp/chromedp v0.7.3
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gin-contrib/static v0.0.1
	github.com/gin-gonic/autotls v0.0.3
	github.com/gin-gonic/gin v1.7.7
	github.com/lifegit/go-gulu/v2 v2.1.7
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.8.1
	github.com/swaggo/gin-swagger v1.3.0
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.2
)

replace gorm.io/datatypes v1.0.6 => github.com/lifegit/datatypes v1.0.7
