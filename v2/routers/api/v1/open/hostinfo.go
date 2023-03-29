package open

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/net"
	"net/http"
	"os/exec"
	"time"
	"zrWorker/core/hostinfo"
	"zrWorker/core/slog"
	"zrWorker/lib/cmd"
	"zrWorker/pkg/utils"
)

var (
	nowtime    string
	hostip     string
	hostinfos  *host.InfoStat
	parts      hostinfo.Parts
	cpuinfos   hostinfo.CpuInfo
	mempercent float64
	meminfos   hostinfo.MemInfo
	netinfos   []net.IOCountersStat
	netspeed   []hostinfo.SpeedInfo

	//nodesstate  []*db.NodeState
	//tasksstate  []db.TaskState

)

// InfoGet 本机信息
func InfoGet(c *gin.Context) {
	var err error
	nowtime = time.Now().Format("2006-01-02 15:04:05")

	cpuinfos, err = hostinfo.GetCpuPercent()
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}
	meminfos = hostinfo.GetMemInfo()
	netinfos, err = hostinfo.GetNetInfo()
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}
	netspeed = hostinfo.GetNetSpeed()

	hostip = hostinfo.GetLocalIP()
	hostinfos, err = hostinfo.GetHostInfo()
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}

	parts, err = hostinfo.GetDiskInfo()
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}
	c.JSON(http.StatusOK,
		gin.H{"hostip": hostip, "hostinfos": hostinfos, "parts": parts,
			"cpuinfos": cpuinfos, "mempercent": mempercent, "meminfos": meminfos,
			"netinfos": netinfos, "netspeed": netspeed, "nowtime": nowtime,
		})
}

// 服务日志
func ServiceLog(c *gin.Context) {

	service := cmd.GetVersion()

	logCmd := exec.Command("tail", "-10", "/zrtx/log/cyberspace/worker.log")

	log, _ := logCmd.CombinedOutput()

	c.JSON(http.StatusOK,
		gin.H{"serviceCmd": string(service), "log": string(log)})
}

// 结果日志
func ResLog(c *gin.Context) {

	logCmd := exec.Command("tail", "-10", "/zrtx/log/cyberspace/ipInfo"+utils.GetHour()+".json")
	log, _ := logCmd.CombinedOutput()

	lCmd := exec.Command("wc", "-l", "/zrtx/log/cyberspace/ipInfo"+utils.GetHour()+".json")
	l, _ := lCmd.CombinedOutput()

	c.JSON(http.StatusOK,
		gin.H{"log": string(log), "count": string(l)})
}
