package main

import (
	"fmt"
	"sync"
)

var wg_A004 sync.WaitGroup

func hello_A04() {
	fmt.Println("Hello World")
	wg_A004.Done()
}

func hello_A04_01() {
	fmt.Println("Hello World 01")
	wg_A004.Done()
}

func hello_A04_02() {
	fmt.Println("Hello World 02")
	wg_A004.Done()
}

func A004() {
	wg_A004.Add(3)

	go hello_A04_01()
	go hello_A04_02()
	hello_A04()

	wg_A004.Wait()

}
