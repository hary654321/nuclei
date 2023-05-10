package cache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {

}

func TestCacheA(t *testing.T) {

	NewCacheClient(time.Duration(1))
	Set("a", []byte("a"))
	println(Get("a"))
}
