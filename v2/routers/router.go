package routers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//r.Use(gin.Logger())
	//pprof.Register(r)
	r.Use(gin.Recovery())
	r.Use(Cors())
	gin.SetMode("debug")
	//setLogFom(r)
	//loadRes(r)

	loadApiV1Open(r)
	return r
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization,Token,X-TOKEN ")
		c.Header("Access-Control-Allow-Methods", "POST, GET,PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func setLogFom(r *gin.Engine) {
	f, _ := os.OpenFile("./app01.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	// 配置中间件
	//r.Use(gin.LoggerWithWriter(io.MultiWriter(f,os.Stdout)))
	// 返回什么格式,日志格式就是什么样子
	var conf = gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("客户端IP:%s,请求时间:[%s],请求方式:%s,请求地址:%s,http协议版本:%s,请求状态码:%d,响应时间:%s,客户端:%s，错误信息:%s\n",
				param.ClientIP,
				param.TimeStamp.Format("2006年01月02日 15:03:04"),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		},
		Output: io.MultiWriter(os.Stdout, f),
	}
	r.Use(gin.LoggerWithConfig(conf))
}

func loadRes(r *gin.Engine) {
	//r.LoadHTMLGlob("dist/*.html")              // 添加入口index.html
	//r.LoadHTMLFiles("./css/*")                   // 添加资源路径
	//r.Static("/css", "./dist/css")             // 添加资源路径
	//r.LoadHTMLFiles("fonts/*")                 // 添加资源路径
	//r.Static("/fonts", "./dist/fonts")         // 添加资源路径
	//r.LoadHTMLFiles("img/*")                   // 添加资源路径
	//r.Static("/img", "./dist/img")             // 添加资源路径
	//r.LoadHTMLFiles("js/*")                    // 添加资源路径
	//r.Static("/js", "./dist/js")               // 添加资源路径
	//r.StaticFile("/", "dist/index.html") //前端接口
}
