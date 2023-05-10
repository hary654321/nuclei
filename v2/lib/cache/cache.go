package cache

import (
	"context"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/projectdiscovery/nuclei/v2/core/slog"
)

var cli *bigcache.BigCache

func NewCacheClient(LifeWindow time.Duration) {

	cli, _ = bigcache.New(context.Background(), bigcache.DefaultConfig(LifeWindow*time.Minute))

}

func Set(k string, v []byte) {
	err := cli.Set(k, v)

	if err != nil {
		slog.Println(slog.DEBUG, "Set err", err.Error())
	}
}

func Get(k string) []byte {
	res, err := cli.Get(k)
	if err != nil {
		//slog.Println(slog.DEBUG, "Get err:", err.Error())
	}

	//slog.Println(slog.INFO, "cache Get info:", k,string(res))

	return res
}
