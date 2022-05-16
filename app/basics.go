package app

import (
	"github.com/lifegit/go-gulu/v2/pkg/logging"
	"github.com/sirupsen/logrus"
	"time"
)

var Log *logrus.Logger

func SetUpBasics() {
	// timeZone
	time.Local, _ = time.LoadLocation(Global.App.TimeZone)

	// log
	Log = logging.NewLogger(Global.App.Log, 3, &logrus.TextFormatter{}, nil)
}
