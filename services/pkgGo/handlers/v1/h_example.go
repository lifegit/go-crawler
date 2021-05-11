package v1

import (
	"github.com/gin-gonic/gin"
	"go-gulu/app"
)

func init() {
	api.GET("example", example)
}

//All
func example(c *gin.Context) {
	app.JsonSuccess(c)
}
