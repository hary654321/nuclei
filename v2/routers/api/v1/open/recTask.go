package open

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/projectdiscovery/nuclei/v2/cmd/nuclei"
)

func RecTask(c *gin.Context) {

	mul := c.PostForm("mul")
	// var ipLastPath = "/zrtx/log/cyberspace/ipLast" + utils.GetHour() + ".json"
	if mul != "" {
		strArrayNew := strings.Split(mul, ",")
		go nuclei.Scan(strArrayNew)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": mul,
	})
}
