package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной
*/

func main() {
	a := 5
	b := 10
	a, b = b, a

	fmt.Println(a, b)
}
