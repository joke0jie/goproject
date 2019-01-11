package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

var mu sync.Mutex
var payloadNum int
var payloadCurNum int
var wg sync.WaitGroup

var payload []string = make([]string, 0)
var pay chan string = make(chan string)
var done bool = true

var client http.Client

//设置请求

func test() {
	resp, err := http.Get("http://218.18.229.57")
	resp.Body.Close()
	if err == nil {

		fmt.Println("1-1-1-1-1--", resp.StatusCode)

	}
}

func main() {
	//test()
	//IISShortscan("123IISShortscan")

	runtime.GOMAXPROCS(runtime.NumCPU())
	t := time.Now()

	target := flag.String("target", "127.0.0.1", "target IP")
	threadNum := flag.Int("thread", 10, "thread num  __note__ port scan port default 15")
	fg := flag.String("flag", "null", `WDS = web dir scan
	useage: ./GScan --flag WDS --target http://127.0.0.1 --thread 50
WPS = web C port scan 80/443/8080 [format 224.221.224]
	usage: ./GScan --flag WPS --target 114.114.114
ISC = IIS short name scan. 
	usage: ./GScan --flag ISC --target http://127.0.0.1/`)

	flag.Parse()
	fmt.Println("[*] cpu:", runtime.NumCPU())
	switch *fg {
	case "WDS":
		{
			WebDirScan(*target, *threadNum)

		}
	case "WPS":
		{
			M_CPortScan(*target)

		}
	case "ISC":
		{
			IISShortscan(*target)
		}
	}

	wg.Wait()
	elapsed := time.Since(t)
	fmt.Println("[*] time:", elapsed)
	fmt.Println("[*] done ...")
}
