package main

import (
	"fmt"
	"strings"
)

func main() {
	content := ("2026-10-31 21:19:29 | 111.55.215.121")
	slic := strings.Split(content, "|")
	for _, v := range slic {
		fmt.Println(v)

	}
}
