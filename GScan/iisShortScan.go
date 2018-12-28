// iisShortScan.go
package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

//1.通过   http://sdl.me/*~1*/.aspx   来判断是否存在短文件名枚举漏洞

//2.如果有 从26个字母中选出一个做文件首字母判断是否存在   /a*~1*/.aspx  404代表存在

//3.如果以a开头的字母存在，那枚举第2个，以此类推直接返回 400 代表枚举完成123

func IISShortscan(target string) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	//设置请求超时时间
	timeout := time.Duration(1000 * time.Millisecond) //超时时间50ms

	//设置请求
	client := &http.Client{Transport: tr, Timeout: timeout}
	client.Get("http://127.0.0.1")

	cha := "qwertyuiopasdfghjklzxcvbnm1234567890"
	for k, v := range cha {

		fmt.Printf("%d - %c\n", k, v)
	}

	fmt.Println()
}
