package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
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

func scan(payload string, target string) {

	if !done {
		return
	}
	baseurl := "http://" + target + payload
	url := strings.Replace(baseurl, "\n", "", -1)
	//fmt.Println("index:", payloadCurNum, url)
	resp, err := http.Get(url)
	if err == nil {
		if resp.StatusCode == 200 {
			fmt.Printf("[[+] \x1b[0;%dm 200       %s \n\x1b[0m", 35, url)

			return
		}
	}
	//判断https
	baseurl = "https://" + target + payload
	url = strings.Replace(baseurl, "\n", "", -1)
	resp, err = http.Get(url)
	if err == nil {
		if resp.StatusCode == 200 {
			fmt.Printf("[[+] \x1b[0;%dm 200       %s \n\x1b[0m", 35, url)
		}
	}

}

func init() {
	payloadFile := "payload.txt"
	payloadNum = 0
	payloadCurNum = 0
	fp, err := os.Open(payloadFile)

	buf := bufio.NewReader(fp)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	for {
		line, err := buf.ReadString('\n')

		if err != nil || err == io.EOF {
			break

		}
		payloadNum += 1
		payload = append(payload, line)
	}
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	//runtime.GOMAXPROCS(1)
	t := time.Now()

	target := flag.String("target", "127.0.0.1", "target IP")
	threadNum := flag.Int("thread", 10, "thread num  __note__ port scan port default 15")
	fg := flag.String("flag", "null", `web = web dir scan
	useage: ./GScan --flag web --target 127.0.0.1 --thread 20
port = web C port scan 80/443/8080 [format 224.221.224]
	usage: ./GScan --flag port --target 114.114.114
ISC = IIS short name scan. 
	usage: ./GScan --target http://127.0.0.1/`)

	flag.Parse()
	//fmt.Println("[+] start ...")
	fmt.Println("[*] cpu:", runtime.NumCPU())
	switch *fg {
	case "web":
		{
			fmt.Printf("[*] target : %s  thread : %d\n", *target, *threadNum)
			fmt.Println("[*] web dir scan start...")
			fmt.Println("[*] payload number : ", payloadNum)
			wg.Add(*threadNum)
			for i := 0; i < *threadNum; i++ {

				go createTask(*target)
			}
		}
	case "port":
		{
			fmt.Printf("[*] target : %s  thread : 15 (default)\n", *target)
			wg.Add(15)
			fmt.Println("[*] port scan start...")
			for i := 0; i < 15; i++ {
				go CportScan(*target, i*17, 16)
			}
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

func CportScan(target string, start int, count int) {
	////设置tls配置信息
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	//设置请求超时时间
	timeout := time.Duration(1000 * time.Millisecond) //超时时间50ms

	//设置请求
	client := &http.Client{Transport: tr, Timeout: timeout}

	for i := start; i <= start+count; i++ {
		IP := target + "." + strconv.Itoa(i)

		uri := "http://" + IP //+ ":80"
		//mt.Println("uri:", uri)
		resp, err := client.Get(uri)
		if err == nil {
			if resp.StatusCode == 200 {
				//fmt.Printf("IP:%s   port:80        ", IP)

				Server := resp.Header.Get("Server")
				//fmt.Printf("Server:%s\t\t\t\t", Server)
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					// handle error
				}
				title := ""
				exp := regexp.MustCompile(`<title>(.*?)</title>`)
				result := exp.FindAllStringSubmatch(string(body), -1)
				for _, text := range result {
					//fmt.Println("result", text[1])
					title = text[1]
				}
				//fmt.Printf("Title:%s       \n", title)
				fmt.Printf("[+] %s\t80\t\t%-20s\t\t\t\t\t%-10s\n", IP, Server, title)

			}

		}

		uri = "https://" + IP
		//fmt.Println("uri:", uri)
		resp, err = client.Get(uri)
		if err == nil {
			if resp.StatusCode == 200 {
				//fmt.Printf("IP:%s   port:443        ", IP)

				Server := resp.Header.Get("Server")
				//fmt.Printf("Server:%s\t\t\t\t", Server)
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {

					// handle error

				}
				title := ""
				exp := regexp.MustCompile(`<title>(.*?)</title>`)
				result := exp.FindAllStringSubmatch(string(body), -1)
				for _, text := range result {
					//fmt.Println("result", text[1])
					title = text[1]
				}
				//fmt.Printf("Title:%s       \n", title)
				fmt.Printf("[+] %s\t443\t\t%-20s\t\t\t\t\t%-10s\n", IP, Server, title)
			}
		}

		uri = "http://" + IP + ":8080"
		resp, err = client.Get(uri)
		if err == nil {
			if resp.StatusCode == 200 {
				//fmt.Printf("IP:%s   port:8080        ", IP)

				Server := resp.Header.Get("Server")
				//fmt.Printf("Server:%s\t\t\t\t", Server)
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {

					// handle error

				}
				title := ""
				exp := regexp.MustCompile(`<title>(.*?)</title>`)
				result := exp.FindAllStringSubmatch(string(body), -1)
				for _, text := range result {
					//fmt.Println("result", text[1])
					title = text[1]
				}
				//fmt.Printf("Title:%s       \n", title)
				fmt.Printf("[+] %s\t8080\t\t%-40s\t\t%-10s\n", IP, Server, title)
			}
		}
	}
	wg.Done()
}

func GetPayload() string {
	mu.Lock()
	if payloadCurNum >= payloadNum {
		done = false
		return ""
	}
	p := payload[payloadCurNum]
	payloadCurNum += 1
	mu.Unlock()
	return p
}

func createTask(target string) {

	for done {
		scan(GetPayload(), target)
	}
	wg.Done()
}
