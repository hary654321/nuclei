package open

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/projectdiscovery/nuclei/v2/cmd/nuclei"
	"github.com/projectdiscovery/nuclei/v2/pkg/utils"
)

func RecTask(c *gin.Context) {

	mul := c.PostForm("mul")
	var ipLastPath = "/zrtx/log/cyberspace/ipLast" + utils.GetHour() + ".json"
	if mul != "" {
		strArrayNew := strings.Split(mul, ",")
		for _, v := range strArrayNew {
			//slog.Println(slog.DEBUG, v)
			utils.WriteAppend(ipLastPath, v)
			go nuclei.Scan(v)
			time.Sleep(1 * time.Second)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": mul,
	})
}
