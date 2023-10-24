package open

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hary654321/nuclei/v2/cmd/nuclei"
	"github.com/hary654321/nuclei/v2/core/slog"
	"github.com/hary654321/nuclei/v2/lib/cmd"

	"github.com/hary654321/nuclei/v2/pkg/utils"
)

func GetZip(c *gin.Context) {

	slog.Printf(slog.DEBUG, "GetZip")
	day := c.Query("day")
	if day == "" {
		day = utils.GetHour()
	}
	target := utils.LogPath + "/" + day + ".zip"

	go utils.ZipFile(utils.LogPath, target, "*"+day+".json")

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": target,
	})
}

func TaskCount(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": nuclei.TaskCount + cmd.TaskCount,
	})
}
