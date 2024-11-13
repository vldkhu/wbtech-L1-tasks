package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной с использованием XOR
*/

func main() {
	a := 5  // Инициализация переменной a со значением 5 (в двоичном виде: 0101)
	b := 10 // Инициализация переменной b со значением 10 (в двоичном виде: 1010)

	// a XOR b
	a ^= b // Теперь a = 5 XOR 10 = 15 (в двоичном виде: 1111)

	// b XOR a
	b ^= a // Теперь b = 10 XOR 15 = 5 (в двоичном виде: 0101)

	// a становится результатом a XOR b
	a ^= b // Теперь a = 15 XOR 5 = 10 (в двоичном виде: 1010)

	fmt.Println(a, b) // Ожидаемый вывод: 10 5
}
