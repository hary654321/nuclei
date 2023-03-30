package main

import (
	"github.com/projectdiscovery/nuclei/v2/lib/cmd"
	"github.com/projectdiscovery/nuclei/v2/pkg/utils"
	"github.com/projectdiscovery/nuclei/v2/routers"
)

func main() {

	cmd.CleanLog()
	utils.Write("/zrtx/log/cyberspace/worker.log", "")
	var restart = "/zrtx/log/cyberspace/restart" + utils.GetHour() + ".json"
	utils.WriteAppend(restart, utils.GetTime())
	r := routers.InitRouter()
	r.Run("0.0.0.0:18000")
}
