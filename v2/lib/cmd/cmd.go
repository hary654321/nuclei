package cmd

import (
	"os"
	"os/exec"
	"time"

	"github.com/projectdiscovery/nuclei/v2/core/slog"
	"github.com/projectdiscovery/nuclei/v2/pkg/utils"
)

func GetVersion() string {

	pid := os.Getpid()

	cmd := exec.Command("ps", "-p", utils.GetInterfaceToString(pid), "-o", "comm=")

	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}

	return string(out)

}

func CleanLog() {

	cmd := exec.Command("bash", "-c", "find /zrtx/log/cyberspace  -mtime +1 -name \"*\" | xargs -I {} rm -rf {}")
	cmd1 := exec.Command("bash", "-c", "rm -rf  /tmp/nuc* && rm -rf /tmp/worker*")
	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}
	out1, err1 := cmd1.CombinedOutput()
	if err1 != nil {
		slog.Println(slog.DEBUG, err1)
	}

	slog.Println(slog.DEBUG, string(out), string(out1))
}

func Dirmap(addr string) {

	cmd := exec.Command("bash", "-c", "cd /tmp && python3  dirmap.py -i "+addr+" -lcf")
	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}

	slog.Println(slog.DEBUG, string(out))
}

var (
	max       = 15
	TaskCount = 0
)

func Scan(ip string) {
	for {
		if TaskCount > max {
			time.Sleep(1 * time.Second)
		} else {
			Dirmap(ip)
			TaskCount -= len(ip)
			slog.Println(slog.DEBUG, "DirmapTaskCount:", TaskCount)
			break
		}
	}

}
