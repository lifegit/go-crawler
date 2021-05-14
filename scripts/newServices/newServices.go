/**
* @Author: TheLife
* @Date: 2021/5/11 上午8:50
 */
package newServices

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path"
)

func Run(gc Bro) error {
	for _, tplNode := range parseOneList {
		err := tplNode.ParseExecute(path.Join(gc.ServiceDir, gc.ServiceName), "", &gc)
		if err != nil {
			return fmt.Errorf("parse [%s] template failed with error : %s", tplNode.NameFormat, err)
		}
	}

	// go fmt codebase
	// https://cloud.tencent.com/developer/article/1417112
	if err := gc.goFmtCodeBase(); err != nil {
		logrus.WithError(err).Error("go fmt code base failed")
	}

	return nil
}

func (app *Bro) goFmtCodeBase() error {
	cmd := exec.Command("go", "fmt", "./...")
	cmd.Dir = path.Join(app.ServiceDir, app.ServiceName)
	cmd.Env = append(os.Environ(), "GOPROXY=https://goproxy.io")
	bb, err := cmd.CombinedOutput()
	if err != nil {
		//print gin-goinc/autols failure
		// fix it :::  https://github.com/gin-gonic/gin/issues/1673
		return fmt.Errorf("%s   %s", string(bb), err)
	}
	return nil
}
