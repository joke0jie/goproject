package main

import (
	_ "bufio"
	"crypto/tls"
	_ "flag"
	"fmt"
	_ "io"
	_ "io/ioutil"
	"net/http"
	_ "os"
	_ "strings"
	_ "sync"
	"time"
)

func main() {

	////设置tls配置信息
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	timeout := time.Duration(1000 * time.Millisecond) //超时时间50ms
	//设置请求
	client := &http.Client{Transport: tr, Timeout: timeout}

	//	//设置请求超时时间

	//	client = &http.Client{
	//
	//	}

	baseurl := "https://118.25.11.115"
	fmt.Println(baseurl)
	resp, err := client.Get(baseurl)
	if err == nil {
		if resp.StatusCode == 200 {
			fmt.Printf("\x1b[0;%dm 200       %s \n\x1b[0m", 35, baseurl)
		}
	} else {
		fmt.Println("err", err)
	}
}
