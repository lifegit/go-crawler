module go-crawler

go 1.15

require (
	github.com/chromedp/chromedp v0.7.2
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gin-contrib/static v0.0.1
	github.com/gin-gonic/autotls v0.0.3
	github.com/gin-gonic/gin v1.7.1
	github.com/mitchellh/mapstructure v1.1.2
	github.com/qiniu/qmgo v0.9.3
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.7.1
	go-gulu v0.0.0-00010101000000-000000000000
	go.mongodb.org/mongo-driver v1.5.1
	google.golang.org/protobuf v1.24.0 // indirect
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.9
)

replace go-gulu => /Users/yxs/GolandProjects/src/go-gulu
