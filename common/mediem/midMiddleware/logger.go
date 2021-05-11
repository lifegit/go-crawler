package midMiddleware

import (
	"encoding/json"
	"fmt"
	"go-crawler/common/mediem"
	"go-gulu/logging"
	"io"
	"os"
	"time"
)

var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

type LogFormatterParams struct {
	TimeStamp   time.Time   `json:"time"`
	ServiceName string      `json:"service_name"`
	Data        mediem.Data `json:"data"`
}

var stdoutLogFormatter = func(param LogFormatterParams) string {
	statusCode, statusColor, statusContent := func() (string, string, interface{}) {
		if param.Data.Err != nil {
			return "fail", yellow, param.Data.Err.Error()
		}
		return "success", green, param.Data.Data
	}()

	return fmt.Sprintf("[MEDIEM %s] %v |%s %s %s \n%v\n",
		param.ServiceName,
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, statusCode, reset,
		statusContent,
	)
}

func NewLoggerMiddleware(isStdout, isWriter bool, serviceName string, writerDir string) (mediem.HandlerFunc, error) {
	var writer io.Writer
	if isWriter {
		w, err := logging.NewRotateIO(writerDir, 5)
		if err != nil {
			return nil, err
		}
		writer = w
	}

	return func(c *mediem.Context) {
		c.Next()

		timestamp := time.Now()

		param := LogFormatterParams{
			TimeStamp:   timestamp,
			ServiceName: serviceName,
			Data:        c.Result,
		}
		bytes, _ := json.Marshal(param)

		if isStdout {
			fmt.Fprint(os.Stdout, stdoutLogFormatter(param))
		}

		if isWriter {
			fmt.Fprint(writer, fmt.Sprintln(string(bytes)))
		}
	}, nil
}
