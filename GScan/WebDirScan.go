// WebDirScan
package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	_ "net"
	"net/http"
	"os"
	_ "strconv"
	"strings"
	"time"
)

func WebDirScan(target string, threadNum int) {
	//test()
	//init data
	if threadNum >= 20 {
		threadNum = 20
	}
	WDSinitdata()
	fmt.Printf("[*] target : %s  thread : %d\n", target, threadNum)
	fmt.Println("[*] web dir scan start...")
	fmt.Println("[*] payload number : ", payloadNum)
	wg.Add(threadNum)
	for i := 0; i < threadNum; i++ {
		go createTask(target)
	}
}

func createTask(target string) {

	for done {
		scan(GetPayload(), target)
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

func WDSinitdata() {
	//payloadFile := "payload.txt"
	payloadFile := "pl.txt"
	payloadNum = 0
	payloadCurNum = 0
	fp, err := os.Open(payloadFile)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	//设置请求超时时间
	timeout := 1 * time.Second //超时时间50ms
	//
	client = http.Client{Transport: tr,
		Timeout: timeout}

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

func scan(payload string, target string) {

	if !done {
		return
	}

	baseurl := target + payload

	url := strings.Replace(baseurl, "\n", "", -1)
	url = strings.Replace(url, "\r", "", -1)
	//mu.Lock()
	resp, err := client.Get(url)
	//mu.Unlock()
	//fmt.Println("index:", payloadCurNum, url)
	if payloadCurNum%200 == 0 {
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
	if err == nil {
		resp.Body.Close()
		if resp.StatusCode == 200 {
			fmt.Printf("[+] \x1b[0;%dm 200       %s \n\x1b[0m", 35, url)
			return
		} else {
			return
		}
	} else {
		//fmt.Println(err)
		return
	}

}
