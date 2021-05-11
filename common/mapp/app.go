/**
* @Author: TheLife
* @Date: 2021/5/8 下午5:35
 */
package mapp

import (
	"go-crawler/common/appLogging"
	"go-crawler/common/conf"
	"time"
)

func init() {
	// timeZone
	_, _ = time.LoadLocation(conf.GetString("app.timeZone"))

	// resultApi
	Result.Setup()
}

//ServerRun start the server
func ServerRun() {
	appLogging.Log.Infof("app: %s , version: %d at start ...", conf.GetString("app.name"), conf.GetInt("app.version"))

	Result.Running()
}

//Close server app
func ServerClose() {

}
