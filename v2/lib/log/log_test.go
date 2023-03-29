package log

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"math"
	"os"
	"runtime"
	"testing"
)

func TestLog(t *testing.T) {

	// 设置日志格式为json格式
	log.SetFormatter(&log.JSONFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	//log.SetLevel(log.WarnLevel)

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

}

func TestShuchu(t *testing.T) {

	writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	log.SetFormatter(&log.JSONFormatter{})
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}

	log.SetOutput(io.MultiWriter(writer3))
	log.Info("info msg")
}

func TestInt(t *testing.T) {
	x := math.Inf(1)
	//println(x)
	switch {
	case x < 0:
		fmt.Println("<")
	case x > 0:
		fmt.Println(">")
	case x == 0:
		fmt.Println("zero")
	default:
		fmt.Println("something else")
	}
}

//func TestJson(t *testing.T) {
//	input := []byte(`{"key1":[{},{"key2":{"key3":[1,2,3]}}]}`)
//
//	// no path, returns entire json
//	root, _ := sonic.Get(input)
//	raw, _ := root.Raw() // == string(input)
//
//	println(raw)
//	// multiple paths
//	root, _ = sonic.Get(input, "key1", 1, "key2")
//	sub, _ := root.Get("key3").Index(2).Int64()
//	println(sub)
//}

func TestContext(t *testing.T) {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					close(dst)
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 实际使用中应该在这里调用 cancel

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			//cancel() // 这里为了使不熟悉 go 的更能明白在这里调用了 cancel()
			break
		}
	}
}

func TestWin(t *testing.T) {
	sysType := runtime.GOOS
	println(sysType)
}
