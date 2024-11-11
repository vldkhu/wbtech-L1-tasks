package main

import (
	"fmt"
	"sync"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
массива, во второй — результат операции x*2, после чего данные из второго
канала должны выводиться в stdout.
*/

func main() {
	var wg sync.WaitGroup

	ch1 := make(chan int)
	ch2 := make(chan int)

	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Горутина для записи в первый канал
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, v := range array {
			ch1 <- v
		}
		close(ch1) // Закрываем канал после записи всех значений
	}()

	// Горутина для обработки данных из первого канала и записи во второй
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch1 {
			ch2 <- v * 2
		}
		close(ch2) // Закрываем второй канал после обработки
	}()

	// Горутина для ожидания завершения всех горутин
	go func() {
		wg.Wait() // Ждем завершения всех горутинн
	}()

	// Чтение из второго канала и вывод в stdout
	for v := range ch2 {
		fmt.Println(v)
	}
}