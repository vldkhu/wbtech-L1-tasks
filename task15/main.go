package main

import (
	"fmt"
	"strings"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации.




var justString string

	func someFunc() {
		v := createHugeString(1 << 10)
		justString = v[:100]
	}

	func main() {
		someFunc()
	}

	Память: Функция createHugeString создаёт очень большую строку (в данном случае размером 1024 байта), и если эта строка не освобождается, то она будет занимать память, даже если мы используем только её часть.
		В Go строки являются неизменяемыми, и когда вы создаёте срез строки, оригинальная строка остаётся в памяти до тех пор, пока на неё есть ссылки.

	Утечка памяти: Если createHugeString создаёт строку, которая занимает много памяти, и вы не используете её полностью, это может привести к утечкам памяти, особенно если someFunc вызывается многократно.

	Производительность: Создание больших строк может быть затратным по времени, и если это происходит в цикле или в многократных вызовах, это может негативно сказаться на производительности приложения
*/

func createHugeString(size int) string {
	// Создаём строку заданного размера и заполняем её символами
	return strings.Repeat("(*_*)", size) // Заполняем строку символом
}

func someFunc() string {
	v := createHugeString(1 << 10) // Создаём строку размером 1024 байта
	return v[:100]                 // Возвращаем только первые 100 символов
}

func main() {
	justString := someFunc() // Получаем результат и сохраняем его
	fmt.Println(justString)  // Используем строку
}