package v1

import (
	"go-crawler/app"
	"go-crawler/services/pkgGo/constant"
)

func Run() {
	v1 := app.NewAweHandlers(constant.ServiceName, "v1")
	{
		v1.GET("example", example)
	}
}
