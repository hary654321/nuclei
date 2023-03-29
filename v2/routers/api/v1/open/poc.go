package open

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zrWorker/nuclei"
	"zrWorker/lib/cmd"
)

func PocScan(c *gin.Context) {

	nuclei.Poc("182.92.154.36","a.json")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": cmd.GetVersion(),
	})
}
