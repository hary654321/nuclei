package open

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"zrWorker/app"
	"zrWorker/core/spy"
	"zrWorker/global"
	"zrWorker/lib/cache"
	"zrWorker/pkg/e"
	"zrWorker/pkg/utils"
	"zrWorker/run"
)

func RecTask(c *gin.Context) {

	t := c.PostForm("t")
	spyParam := c.PostForm("spy")
	mul := c.PostForm("mul")
	hydra := c.PostForm("hydra")
	addr := c.PostForm("addr")
	var ipLastPath = "/zrtx/log/cyberspace/ipLast" + utils.GetHour() + ".json"
	if t != "" {
		global.AppSetting.Target = append(global.AppSetting.Target, t)
		run.PushTarget(t)
	}
	if spyParam != "" {
		spy.Start(spyParam)
	}
	if mul != "" {
		strArrayNew := strings.Split(mul, ",")
		for _, v := range strArrayNew {
			//slog.Println(slog.DEBUG, v)
			utils.WriteAppend(ipLastPath, v)
			run.PushTarget(v)
		}
	}
	if hydra == "1" {
		app.Setting.Hydra = true
	}
	if addr != "" {
		//global.AppSetting.AdrrArr = append(global.AppSetting.AdrrArr, addr)
		//go func() {
		//	run.ScanAddr(addr)
		//}()
	}
	//任务信息记录下来
	startTime := utils.GetTime()
	runTaskID := c.PostForm("runTaskId")
	taskId := c.PostForm("taskId")
	logData := cache.TaskLog{
		TaskID:    taskId,
		RunTaskID: runTaskID,
		StartTime: startTime,
		Progress:  0,
	}

	cache.SetTaskLog(runTaskID, logData)

	data := make(map[string]interface{})
	code := e.SUCCESS
	data["taskId"] = taskId
	data["runTaskId"] = runTaskID
	data["startTime"] = startTime

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
