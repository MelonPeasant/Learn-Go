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

func hello_A04_01(i int) {
	fmt.Println("Hello World 01", i)
	wg_A004.Done()
}

func hello_A04_02() {
	fmt.Println("Hello World 02")
	wg_A004.Done()
}

func A004() {
	wg_A004.Add(12)
	for i := 0; i < 10; i++ {
		go hello_A04_01(i)
	}
	go hello_A04_02()
	hello_A04()

	wg_A004.Wait()

}
