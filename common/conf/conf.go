package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

const path = "./conf"
const confType = "toml"
const confMain = "base"
const confNormal = "conf"

func init() {
	watch := handWatchFileChange
	// mainConf init
	v, err := getLocalConfToViper(path, confMain, confType, &watch)
	if err != nil {
		logrus.WithError(err).Fatal(err)
		return
	}
	// setting to default
	AllSettingsToDefault(v)

	// check ENV
	//env := os.Getenv("GO_ENV")
	v, err = getLocalConfToViper(fmt.Sprintf("%s/%s/", path, viper.GetString("app.env")), confNormal, confType, &watch)
	if err != nil {
		logrus.WithError(err).Fatal(err)
		return
	}
	// setting to default
	AllSettingsToDefault(v)

	//fmt.Println("local","AllSettings",viper.AllSettings())
	//return nil
}
func AllSettingsToDefault(setting *viper.Viper) {
	configs := setting.AllSettings()
	for k, v := range configs {
		viper.SetDefault(k, v)
	}
}

// handWatchFileChange
func handWatchFileChange(event fsnotify.Event) {
	if event.Op == fsnotify.Create || event.Op == fsnotify.Write || event.Op == fsnotify.Chmod {
		path := event.Name[:strings.LastIndex(event.Name, "/")]
		pathname := event.Name[strings.LastIndex(event.Name, "/")+1 : strings.LastIndex(event.Name, ".")]
		confType := event.Name[strings.LastIndex(event.Name, ".")+1:]
		v, err := getLocalConfToViper(path, pathname, confType, nil)
		if err == nil {
			// setting to default
			AllSettingsToDefault(v)
			_ = fmt.Sprintf("application configuration'initialization watch success in %s", event.Name)
			return
		}
	}

	_ = fmt.Sprintf("application configuration'initialization watch fail in %s, file op is %s", event.Name, event.Op)
}

// getFileConfToViper
func getLocalConfToViper(path, pathname, confType string, WatchChange *func(fsnotify.Event)) (*viper.Viper, error) {
	// viper init
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigType(confType)
	v.SetConfigName(pathname)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	if WatchChange != nil {
		v.WatchConfig()
		v.OnConfigChange(*WatchChange)
	}
	return v, nil
}

// GetString returns the value associated with the key as a string.
func GetString(key string) string {
	return viper.GetString(key)
}

// GetInt returns the value associated with the key as an integer.
func GetInt(key string) int {
	return viper.GetInt(key)
}

// GetBool returns the value associated with the key as a boolean.
func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}
