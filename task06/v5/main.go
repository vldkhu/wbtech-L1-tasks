package main

import (
	"context"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
_______________________________________________________________

Четвертый способ это прерывание по таймауту. контекст с таймаутом
*/

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // Проверяем, истек ли таймаут
				return // Завершаем горутину
			default:
				// Выполняем основную работу
			}
		}
	}(ctx)
}
