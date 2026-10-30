package main

import (
	"flag"
	"fmt"
)

func A005() {
	var name_A005 string
	var port_A005, age_A005 int

	flag.StringVar(&name_A005, "name", "NONE", "名字")
	flag.IntVar(&port_A005, "port", 22, "端口")
	flag.IntVar(&age_A005, "age", 35, "年龄")

	flag.Parse()

	fmt.Println(name_A005)
	fmt.Println(port_A005)
	fmt.Println(age_A005)
}
