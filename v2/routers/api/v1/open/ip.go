package open

import (
	"github.com/gin-gonic/gin"
	"net/http"
	ip2 "zrWorker/lib/ip"
	"zrWorker/pkg/e"
)

func Ip(c *gin.Context) {

	ip := c.Query("ip")

	res := ip2.SearchIpAddr(ip)

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  "e.GetMsg(code)",
		"data": res,
	})
}
