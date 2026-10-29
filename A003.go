package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Work(i int) {
	fmt.Printf("第%d个Work正在工作...\n", i)
	wg.Done()
}

func A003() {
	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go Work(i)
	}
	wg.Wait()
}
