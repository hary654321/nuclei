package open

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zrWorker/lib/cmd"
	"zrWorker/pkg/utils"
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
