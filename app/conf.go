package app

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lifegit/go-gulu/v2/nice/file"
	"github.com/lifegit/go-gulu/v2/pkg/viperine"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
)

var Global GlobalConf

type GlobalConf struct {
	App struct {
		Name     string `toml:"name"`
		Version  int    `toml:"version"`
		TimeZone string `toml:"timeZone"`
		Env      string `toml:"env"`
		Log      string `toml:"log"`
	} `toml:"app"`
	Server struct {
		Addr         string `toml:"addr"`
		Port         int    `toml:"port"`
		RunMode      string `toml:"runMode"`
		ReadTimeout  int    `toml:"readTimeout"`
		WriteTimeout int    `toml:"writeTimeout"`
		APIPrefix    string `toml:"apiPrefix"`
		StaticPath   string `toml:"staticPath"`
		HTTPLogDir   string `toml:"httpLogDir"`
		IsNotRoute   bool   `toml:"isNotRoute"`
		IsSwagger    bool   `toml:"isSwagger"`
		IsCors       bool   `toml:"isCors"`
		IsHTTPS      bool   `toml:"isHttps"`
	} `toml:"server"`
	Service struct {
		Log string `toml:"log"`
	} `toml:"service"`
	Db struct {
		Type     string `toml:"string"`
		Addr     string `toml:"addr"`
		Port     int    `toml:"port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
		Database string `toml:"database"`
		Charset  string `toml:"charset"`
	} `toml:"db"`
}

const DEV = "dev"

func (g *GlobalConf) isDev() bool {
	return g.getEnv() == DEV
}
func (g *GlobalConf) getEnv() (res string) {
	if res = os.Getenv("GO_ENV"); res == "" {
		res = DEV
	}

	return res
}
func SetUpConf() {
	basePath := recursionPath("conf")
	v, err := viperine.LocalConfToViper([]string{
		path.Join(basePath, "base.toml"),
		path.Join(basePath, Global.getEnv(), "conf.toml"),
	}, &Global, func(event fsnotify.Event, viper *viper.Viper) {
		if event.Op != fsnotify.Remove {
			_ = viper.Unmarshal(&Global)
		}
	})

	if err != nil {
		logrus.WithError(err).Fatal(err, v)
	}
}

func recursionPath(dirName string) (dirPath string) {
	var dir string
	for i := 0; i < 10; i++ {
		dirPath = path.Join(dir, dirName)
		dir = path.Join(dir, "../")

		if file.IsDir(dirPath) {
			return
		}
	}

	return
}
