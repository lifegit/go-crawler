package v1

import (
	"github.com/gin-gonic/gin"
	"go-gulu/app"
)

//All
func example(c *gin.Context) {
	app.JsonSuccess(c)
}
