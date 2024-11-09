package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Map // Создаем sync.Map для безопасного доступа

	start := time.Now() // Запоминаем время начала

	// Заполняем sync.Map значениями от 0 до 9999
	for i := 0; i < 10_000; i++ {
		m.Store(i, i) // Сохраняем ключ и значение
	}

	// Выводим все элементы из sync.Map
	m.Range(func(key, value any) bool {
		fmt.Println(key, value) // Печатаем ключ и значение
		return true             // Продолжаем итерацию
	})

	elapsed := time.Since(start)              // Вычисляем время выполнения
	fmt.Printf("Время чтения: %s\n", elapsed) // Выводим время
}
