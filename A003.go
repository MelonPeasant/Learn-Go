package main

import (
	"fmt"
	"sync"
)

var wg_A003 sync.WaitGroup

func Work(i int) {
	fmt.Printf("第%d个Work正在工作...\n", i)
	wg_A003.Done()
}

func A003() {
	wg_A003.Add(5)
	for i := 1; i <= 5; i++ {
		go Work(i)
	}
	wg_A003.Wait()
}
