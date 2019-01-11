// CPortScan
package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var Portvalue int

func M_CPortScan(target string) {
	Portvalue = 0
	fmt.Printf("[*] target : %s  thread : 15 (default)\n", target)
	wg.Add(15)
	fmt.Println("[*] port scan start...")
	for i := 0; i < 15; i++ {
		go CportScan(target, i*17, 16)
	}
	//fmt.Println("[+] count:", Portvalue)
}

func printResult(resp *http.Response, IP string, port int, code int) {
	Server := resp.Header.Get("Server")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	title := ""
	exp := regexp.MustCompile(`<title>(.*?)</title>`)
	result := exp.FindAllStringSubmatch(string(body), -1)
	for _, text := range result {
		title = text[1]
	}
	fmt.Printf("[+] %s:%d\t%d\t\t%-40s\t\t%-10s\n", IP, port, code, Server, title)

	Portvalue += 1

}

func CportScan(target string, start int, count int) {
	////设置tls配置信息
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	//设置请求超时时间
	timeout := time.Duration(1500 * time.Millisecond) //超时时间50ms

	//设置请求
	client := &http.Client{Transport: tr, Timeout: timeout}

	for i := start; i <= start+count; i++ {
		IP := target + "." + strconv.Itoa(i)

		uri := "http://" + IP
		resp, err := client.Get(uri)

		//		if i == 57 {
		//			if err != nil {
		//				fmt.Println(err)
		//			} else {
		//				fmt.Println("----", resp.StatusCode)
		//			}

		//		}

		if err == nil {
			defer resp.Body.Close()
			switch resp.StatusCode {
			case 200:
				printResult(resp, IP, 80, 200)
				break
			case 404:
				printResult(resp, IP, 80, 404)
				break
			case 302:
				printResult(resp, IP, 80, 302)
				break
			}

		}

		uri = "https://" + IP
		resp, err = client.Get(uri)
		if err == nil {
			defer resp.Body.Close()
			switch resp.StatusCode {
			case 200:
				printResult(resp, IP, 443, 200)
				break
			case 404:
				printResult(resp, IP, 443, 404)
				break
			case 302:
				printResult(resp, IP, 443, 302)
				break
			}
		}

		uri = "http://" + IP + ":8080"
		resp, err = client.Get(uri)
		if err == nil {
			defer resp.Body.Close()
			switch resp.StatusCode {
			case 200:
				printResult(resp, IP, 8080, 200)
				break
			case 404:
				printResult(resp, IP, 8080, 404)
				break
			case 302:
				printResult(resp, IP, 8080, 302)
				break
			}
		}

		uri = "http://" + IP + ":7001"
		resp, err = client.Get(uri)
		if err == nil {
			defer resp.Body.Close()
			switch resp.StatusCode {
			case 200:
				printResult(resp, IP, 7001, 200)
				break
			case 404:
				printResult(resp, IP, 7001, 404)
				break
			case 302:
				printResult(resp, IP, 7001, 302)
				break
			}
		}

		uri = "http://" + IP + "8081"
		resp, err = client.Get(uri)
		if err == nil {
			defer resp.Body.Close()
			switch resp.StatusCode {
			case 200:
				printResult(resp, IP, 8081, 200)
				break
			case 404:
				printResult(resp, IP, 8081, 404)
				break
			case 302:
				printResult(resp, IP, 8081, 302)
				break
			}
		}

		uri = "http://" + IP + ":8082"
		resp, err = client.Get(uri)
		if err == nil {
			defer resp.Body.Close()
			switch resp.StatusCode {
			case 200:
				printResult(resp, IP, 8082, 200)
				break
			case 404:
				printResult(resp, IP, 8082, 404)
				break
			case 302:
				printResult(resp, IP, 8082, 302)
				break
			}
		}
	}
	wg.Done()
}
