package main

import (
	"fmt"
	"strings"
)

func main() {
	content := ("2026-10-31 21:19:29 | 111.55.215.121")
	slic := strings.Split(content, "|")
	for i, v := range slic {
		slic[i] = strings.TrimSpace(v)

		fmt.Println(strings.TrimSpace(v))
	}
	fmt.Println(slic)

}
