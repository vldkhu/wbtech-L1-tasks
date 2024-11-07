package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
// квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	array := [1000000]int{} //добавил миллион, чтобы показать явную эффективность конкуретного подхода.
	for i := 0; i < 1_000_000; i++ {
		array[i] = i + 1
	}
	count := 0

	for _, v := range array {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			sqrt := val * val
			mu.Lock()
			count += sqrt
			mu.Unlock()
		}(v)
	}
	wg.Wait()

	fmt.Println(count)
}
