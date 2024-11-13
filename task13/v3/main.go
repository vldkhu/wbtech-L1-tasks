package main

import (
	"fmt"
	"reflect"
)

/*
Поменять местами два числа без создания временной переменной. Пакет reflect
*/

func main() {
	aSlice := []int{5, 10}

	swap := reflect.Swapper(aSlice)

	swap(0, 1)
	fmt.Println(aSlice[0])
	fmt.Println(aSlice[1])
}
