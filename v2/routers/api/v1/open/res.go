package open

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"zrWorker/app"
	"zrWorker/core/slog"
	"zrWorker/global"
	"zrWorker/lib/cache"
	"zrWorker/pkg/e"
	"zrWorker/pkg/utils"
	"zrWorker/run"
)

func GetZip(c *gin.Context) {

	slog.Printf(slog.DEBUG, "GetZip")
	day := c.Query("day")
	if day == "" {
		day = utils.GetHour()
	}
	target := global.ServerSetting.LogPath + "/" + day + ".zip"

	go utils.ZipFile(global.ServerSetting.LogPath, target, "*"+day+".json")

	//获取文件的名称
	//fileName := path.Base(target)
	//c.Header("Content-Type", "application/octet-stream")
	//c.Header("Content-Disposition", "attachment; filename="+fileName)
	//c.Header("Content-Transfer-Encoding", "binary")
	//c.Header("Cache-Control", "no-cache")
	//c.Header("Content-Type", "application/octet-stream")
	//c.Header("Content-Disposition", "attachment; filename="+fileName)
	//c.Header("Content-Transfer-Encoding", "binary")
	//c.File(target)

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": target,
	})
}

func GetTaskRes(c *gin.Context) {
	runTaskId := c.Query("runTaskId")
	taskInfo := cache.GetTaskLog(runTaskId)

	pCount := run.GetTaskCount()

	taskInfo.Progress = pCount
	taskInfo.Res = utils.Read(utils.GetLogPath("ipInfo"))
	utils.PrinfI("res", taskInfo.Res)
	delete(app.Setting.TaskPercent, runTaskId)

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": taskInfo,
	})
}

// 虚拟进度
func getPercent(runTaskId string, threads int) int {

	if threads == 0 {
		app.Setting.TaskPercent[runTaskId] = 100
	}

	if threads > 0 && threads <= 33 && app.Setting.TaskPercent[runTaskId] < 10 {
		app.Setting.TaskPercent[runTaskId] = 10
	}

	if threads > 33 && threads <= 66 && app.Setting.TaskPercent[runTaskId] < 30 {
		app.Setting.TaskPercent[runTaskId] = 30
	}

	if threads > 66 && threads <= 99 && app.Setting.TaskPercent[runTaskId] < 50 {
		app.Setting.TaskPercent[runTaskId] = 50
	}

	if threads > 99 && app.Setting.TaskPercent[runTaskId] < 80 {
		app.Setting.TaskPercent[runTaskId] = 80
	}

	return app.Setting.TaskPercent[runTaskId]
}

func Image(c *gin.Context) {

	url := c.Query("url")

	filePath := utils.GetScreenPath() + utils.Md5(url) + ".png"

	slog.Printf(slog.DEBUG, filePath)
	//获取文件的名称
	fileName := path.Base(filePath)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(filePath)

	return
}

func TaskCount(c *gin.Context) {

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": run.GetTaskCount(),
	})
}
