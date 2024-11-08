package main

import "context"

/*
Реализовать все возможные способы остановки выполнения горутины.
_______________________________________________________________

Второй способ это использование контекста.
*/

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // Проверяем, был ли отменен контекст
				return // Завершаем горутину
			default:
				// Выполняем основную работу
			}
		}
	}(ctx)

	cancel() // Вызываем cancel для отмены контекста
}
