package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговое
значение счетчика.

2ой способ: с использованием атомика
*/

type Counter struct {
	value int32
}

func main() {
	var wg sync.WaitGroup
	var counter Counter

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 5; j++ {
				atomic.AddInt32(&counter.value, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter.value)
}
