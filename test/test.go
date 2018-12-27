package test

import (
	"fmt"
	"net/http"
)

func main()  {

	url := "www.z-gelen.com"

	fmt.Println("sdff2f2f")
	resp, err := http.Get(url)
	if err == nil {
		if resp.StatusCode == 200 {
			fmt.Printf("\x1b[0;%dm 200       %s \n\x1b[0m", 35, url)

		}
	}else {
		fmt.Println("exit")
	}
}