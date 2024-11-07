/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при
старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
способ завершения работы всех воркеров.
*/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	min = 0
	max = 100
)

func worker(id int, vals <-chan int, wg *sync.WaitGroup, signal chan os.Signal) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении горутины
	for {
		select {
		case val, ok := <-vals:
			if !ok {
				fmt.Printf("goroutine id=%d: channel closed\n", id)
				return
			}
			fmt.Printf("goroutine id=%d value=%d\n", id, val)

		case <-signal:
			fmt.Printf("goroutine id=%d stopped\n", id)
			return
		}
	}
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var amountGo int
	_, err := fmt.Scan(&amountGo)
	if err != nil {
		return
	}

	values := make(chan int, amountGo)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup // Создаем WaitGroup для отслеживания горутин

	// Запуск рабочих горутин
	for i := 0; i < amountGo; i++ {
		wg.Add(1) // Увеличиваем счетчик WaitGroup
		go worker(i, values, &wg, sigChan)
	}

	// Основной цикл для генерации значений
	go func() {
		<-sigChan
		fmt.Printf("main thread is stopped\n")
		close(values) // Закрываем канал после получения сигнала
	}()

	// Генерация значений
	for {
		select {
		case <-sigChan:
			close(values) // Закрываем канал при получении сигнала
			wg.Wait()     // Ждем завершения всех горутин
			return        // Завершаем main, если получен сигнал
		default:
			values <- rand.Intn(max-min+1) + min
			time.Sleep(time.Second / 4)
		}

	}

}
