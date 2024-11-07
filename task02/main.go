package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	var wg sync.WaitGroup

	array := [5]int{2, 4, 6, 8, 10}

	for _, v := range array {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			sqrt := v * v
			fmt.Println(sqrt)
		}(v)
	}
	wg.Wait()
}
