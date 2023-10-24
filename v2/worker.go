package main

import (
	"time"

	"github.com/hary654321/nuclei/v2/cmd/nuclei"
	"github.com/hary654321/nuclei/v2/core/slog"
	"github.com/hary654321/nuclei/v2/lib/cache"
	"github.com/hary654321/nuclei/v2/lib/cmd"
	"github.com/hary654321/nuclei/v2/pkg/utils"
	"github.com/hary654321/nuclei/v2/routers"
)

func main() {

	utils.Write(utils.LogPath+"worker.log", "")
	var restart = utils.LogPath + "restart" + utils.GetHour() + ".json"
	utils.WriteAppend(restart, utils.GetTime())

	cache.NewCacheClient(time.Duration(30))

	r := routers.InitRouter()
	go heart()
	r.Run("0.0.0.0:18000")
}

func heart() {

	for {
		slog.Println(slog.DEBUG, "poc====", nuclei.TaskCount)

		if time.Now().Unix()%3600 < 6 {
			cmd.ResTart()
		}
		time.Sleep(5 * time.Second)
	}
}
