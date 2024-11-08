package main

import (
	"fmt"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
_______________________________________________________________

Первый способ остановки горутин и самый распространенный - это использование каналов
*/

func main() {
	stop := make(chan struct{}) // создаем канал для остановки
	defer close(stop)           // закрываем канал с сигнализации о завершении

	go func() {
		for {
			select {
			case <-stop: // если получен сигнал остановки
				fmt.Println("остановка горутин")
				return // завершаем работу
			default:
				fmt.Println("продолжаем работу") //выполняем основную работу
			}
		}
	}()

}
