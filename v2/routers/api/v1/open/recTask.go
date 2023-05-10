package open

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/projectdiscovery/nuclei/v2/cmd/nuclei"
	"github.com/projectdiscovery/nuclei/v2/lib/cache"
	"github.com/projectdiscovery/nuclei/v2/pkg/utils"
)

func RecTask(c *gin.Context) {

	target := c.PostForm("target")
	tempName := c.PostForm("tempName")
	tempContent := c.PostForm("tempContent")
	tempPath := c.PostForm("tempPath")

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

	temPre := "/root/nuclei-templates"

	tmp := ""
	if tempContent != "" {
		tmp = temPre + "/diy/" + tempName + ".yaml"
		utils.Write(tmp, tempContent)
	} else {
		tmp = temPre + tempPath

		if !utils.FileExists(tmp) {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "模板不存在",
				"data": "",
			})
			return
		}
	}

	if target != "" {
		strArrayNew := strings.Split(target, ",")
		nuclei.Scan(strArrayNew, tmp, taskId)
	}

	data := make(map[string]interface{})

	data["taskId"] = taskId

	data["startTime"] = utils.GetTime()

	data["res"] = utils.Read("/zrtx/log/cyberspace/" + taskId + ".json")

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": data,
	})
}
