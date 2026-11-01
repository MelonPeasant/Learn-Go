package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const FilePath = "ip.txt"

func GetLastLine() (string, bool) {
	file, err := os.Open(FilePath)
	if err != nil {
		// fmt.Println("文件打开错误",err)
		return "文件打开错误", false
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "文件读入错误", false
	}
	if len(content) == 0 {
		return "文件内容为空", false
	}

	slic := strings.Split(strings.TrimSpace(string(content)), "\n")
	for i, v := range slic {
		slic[i] = strings.TrimSpace(v)
	}
	fmt.Println(slic)
	fmt.Println(len(slic))
	res := slic[len(slic)-1]
	return res, true
}

func main() {
	res, status := GetLastLine()
	if status != true {
		fmt.Println(res)
		return
	}
	fmt.Println(res)

}
