package main

import (
	"github.com/projectdiscovery/nuclei/v2/routers"
)

func main() {

	r := routers.InitRouter()
	r.Run("0.0.0.0:18000")
}
