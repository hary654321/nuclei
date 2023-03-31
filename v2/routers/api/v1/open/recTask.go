package open

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/projectdiscovery/nuclei/v2/cmd/nuclei"
	"github.com/projectdiscovery/nuclei/v2/pkg/utils"
)

func RecTask(c *gin.Context) {

	mul := c.PostForm("mul")
	var ipLastPath = "/zrtx/log/cyberspace/ipLast" + utils.GetHour() + ".json"
	if mul != "" {
		strArrayNew := strings.Split(mul, ",")
		go nuclei.Scan(strArrayNew)
		for _, v := range strArrayNew {
			utils.WriteAppend(ipLastPath, v)
		}
	}

	data := make(map[string]interface{})

	data["taskId"] = 1
	data["runTaskId"] = 1
	data["startTime"] = utils.GetTime()

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": data,
	})
}
