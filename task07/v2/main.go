package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Map

	start := time.Now()

	for i := 0; i < 10_000; i++ {
		m.Store(i, i)
	}

	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})

	elapsed := time.Since(start)
	fmt.Printf("Время чтения: %s\n", elapsed)

}
