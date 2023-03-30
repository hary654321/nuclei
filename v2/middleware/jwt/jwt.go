package jwt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/projectdiscovery/nuclei/v2/pkg/utils"
)

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
	INVALID_PASS   = 402
	INVALID_FINGER = 405

	INVALID_DIFFPASS = 403

	ERROR_CRON_SPEC = 10001
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = SUCCESS
		token := c.GetHeader("Authorization")

		if token == "" {
			code = INVALID_PARAMS
		} else {
			claims, err := utils.ParseToken(token)
			if claims == nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": 401,
					"msg":  "cookie失效，请点击右上角退出重新登陆",
					"data": data,
				})
				c.Abort()
				return
			}
			if err != nil {
				code = ERROR
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = ERROR
			}
		}

		if code != SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

func Open() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		token := c.GetHeader("Authorization")
		if token != "Basic enc6Y2g=" {
			code = ERROR
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}

var MsgFlags = map[int]string{
	SUCCESS:          "请求成功",
	ERROR:            "请求失败",
	INVALID_PARAMS:   "请求参数错误",
	INVALID_PASS:     "旧密码错误",
	INVALID_DIFFPASS: "两次新密码不一致",
	INVALID_FINGER:   "需要去掉指纹尾部逗号",

	ERROR_CRON_SPEC: "crontab语法错误",
}
