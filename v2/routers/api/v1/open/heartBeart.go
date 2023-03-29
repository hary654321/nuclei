package open

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/projectdiscovery/nuclei/v2/lib/cmd"
	"github.com/projectdiscovery/nuclei/v2/pkg/utils"
)

type HeartBeatS struct {
	time         string
	version      string
	runningTasks string
}

func HeartBeat(c *gin.Context) {

	data := make(map[string]interface{})
	data["time"] = utils.GetTime()
	data["version"] = cmd.GetVersion()
	data["runningTasks"] = "3412341234"
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "",
		"data": data,
	})
}
