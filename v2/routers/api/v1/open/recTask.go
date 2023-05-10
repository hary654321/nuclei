package open

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/projectdiscovery/nuclei/v2/cmd/nuclei"
	"github.com/projectdiscovery/nuclei/v2/lib/cache"
	"github.com/projectdiscovery/nuclei/v2/lib/cmd"
	"github.com/projectdiscovery/nuclei/v2/pkg/utils"
)

func RecTask(c *gin.Context) {

	target := c.PostForm("target")
	tempName := c.PostForm("tempName")
	tempContent := c.PostForm("tempContent")
	httpAddr := c.PostForm("httpAddr")

	taskId := c.PostForm("taskId")

	if string(cache.Get(taskId)) != "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "任务已存在",
			"data": "",
		})
		return
	}

	cache.Set(taskId, []byte("1"))
	tmp := ""
	if tempName != "" {
		tmp = "/nuclei-templates/diy/" + tempName + ".yaml"
	}

	utils.Write(tmp, tempContent)

	var ipLastPath = "/zrtx/log/cyberspace/ipLast" + utils.GetHour() + ".json"
	if target != "" {
		strArrayNew := strings.Split(target, ",")
		go nuclei.Scan(strArrayNew, tmp, taskId)
		for _, v := range strArrayNew {
			utils.WriteAppend(ipLastPath, v)
		}
	}

	if httpAddr != "" {
		strArrayNew := strings.Split(httpAddr, ",")
		for _, v := range strArrayNew {
			go cmd.Scan(v)
		}
	}

	data := make(map[string]interface{})

	data["taskId"] = taskId

	data["startTime"] = utils.GetTime()

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": data,
	})
}
