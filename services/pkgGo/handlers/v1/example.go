package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lifegit/go-gulu/v2/pkg/out"
)

//All
func example(c *gin.Context) {
	out.JsonSuccess(c)
}
