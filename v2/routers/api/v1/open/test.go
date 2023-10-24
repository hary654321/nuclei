package open

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hary654321/nuclei/v2/lib/cmd"
)

func Test(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": cmd.GetVersion(),
	})
}
