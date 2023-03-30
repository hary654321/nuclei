package routers

import (
	"github.com/projectdiscovery/nuclei/v2/routers/api/v1/open"

	"github.com/gin-gonic/gin"
	"github.com/projectdiscovery/nuclei/v2/middleware/jwt"
)

func loadApiV1Open(r *gin.Engine) {

	openApi := r.Group("/api/v1")
	//开放接口
	openApi.Use(jwt.Open())
	{
		//心跳
		openApi.GET("/heartbeat", open.HeartBeat)

		//接收任务
		openApi.POST("/recTask", open.RecTask)

		//本机信息
		openApi.GET("/info", open.InfoGet)

		//zip压缩包结果
		openApi.GET("/zip", open.GetZip)

		//结果数据
		openApi.GET("/taskCount", open.TaskCount)

		openApi.GET("/serviceLog", open.ServiceLog)

		openApi.GET("/resLog", open.ResLog)

		openApi.GET("/test", open.Test)

	}
}
