package main

import (
	_ "bufio"
	"fmt"
	_ "io"
	"os"
	_ "strconv"
	_ "strings"
)

func main() {
	name := "libva++.so"
	soinfo := new(SoInfo)

	if fileObj, err := os.Open(name); err == nil {
		defer fileObj.Close()
		//解析文件头
		parseHeader(fileObj, soinfo)
		//解析节区
		parseSection(fileObj, soinfo)

		fmt.Println("————————————————————————————————————-")

	}

}
