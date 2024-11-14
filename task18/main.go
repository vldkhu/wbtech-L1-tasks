package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговое
значение счетчика.
	первый способ: с использованием мьютекса
*/

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func main() {
	var counter Counter
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 5; j++ {
				counter.increment()
			}
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println(counter.value)
}
