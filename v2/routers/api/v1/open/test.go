package open

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zrWorker/lib/cmd"
)

func Test(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": cmd.GetVersion(),
	})
}
