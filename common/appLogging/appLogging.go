/**
* @Author: TheLife
* @Date: 2020-2-27 12:12 上午
 */
package appLogging

import (
	"github.com/sirupsen/logrus"
	"go-crawler/common/conf"
	"go-gulu/logging"
)

var Log *logrus.Logger

func init() {
	Log = logging.NewLogger(conf.GetString("app.log"), 3, &logrus.JSONFormatter{}, nil)
}
