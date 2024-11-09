package main

import (
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
_______________________________________________________________

Четвертый способ использования "ok" с каналами
*/

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch) // Закрываем канал после отправки значений
	}()

	for {
		value, ok := <-ch // Читаем из канала
		if !ok {          // Проверяем, закрыт ли канал
			fmt.Println("Канал закрыт, работа заверешена!") // Выходим из цикла, если канал закрыт
			break
		}
		fmt.Println("Получено значение: ", value)
	}

}