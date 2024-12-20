package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Реализовать конкурентную запись данных в map

Первый способ, через мбютекс
*/

// Numbers структура, которая содержит мьютекс для синхронизации доступа к карте значений.
type Numbers struct {
	mu  sync.Mutex  // Мьютекс для защиты доступа к карте
	val map[int]int // Карта для хранения пар "ключ-значение", где ключ и значение - целые числа
}

// Add метод, который добавляет число в карту, обеспечивая безопасный доступ с помощью мьютекса.
func (v *Numbers) Add(num int) {
	v.mu.Lock()         // Блокируем мьютекс для безопасного доступа к карте
	defer v.mu.Unlock() // Обеспечиваем разблокировку мьютекса в конце функции
	v.val[num] = num    // Добавляем число в карту
}

func main() {
	var wg sync.WaitGroup // Создаем WaitGroup для ожидания завершения всех горутин

	start := time.Now()
	// Инициализируем структуру Numbers с пустой картой
	nums := &Numbers{
		val: make(map[int]int, 10_000), // Создаем карту с начальной емкостью 1,000,000
	}

	// Запускаем 1,000,000 горутин для добавления чисел в карту
	for i := 0; i < 10_000; i++ {
		wg.Add(1) // Увеличиваем счетчик WaitGroup на 1 для каждой новой горутины
		go func(nums *Numbers, val int) {
			defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении горутины
			nums.Add(val)   // Вызываем метод Add для добавления числа в карту
		}(nums, i) // Передаем текущую итерацию и указатель на nums в горутину
	}

	wg.Wait() // Ожидаем завершения всех горутин

	// Печатаем содержимое мапы
	fmt.Println(nums.val)

	elapsed := time.Since(start)
	fmt.Printf("Время чтения: %s\n", elapsed)
	// Время чтения: 185.4675ms
}
