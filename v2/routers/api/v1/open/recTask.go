package open

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecTask(c *gin.Context) {

	t := c.PostForm("t")

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": t,
	})
}
