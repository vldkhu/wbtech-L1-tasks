package main

import (
	"fmt"
)

func setBit(n int64, i uint, value int) int64 {
	if value == 1 {
		return n | (1 << i)
	} else {
		return n &^ (1 << i)
	}
}

func main() {
	n := int64(4)  // Исходное значение (в двоичном: 100)
	var i uint = 3 // Индекс бита, который нужно изменить

	// Установка i-го бита в 1
	n = setBit(n, i, 1)
	fmt.Printf("После установки %d-го бита в 1: %b\n", i, n) // Ожидается: 1100 (12 в десятичной системе)

	// Установка i-го бита в 0
	n = setBit(n, i, 0)
	fmt.Printf("После установки %d-го бита в 0: %b\n", i, n) // Ожидается: 100 (4 в десятичной системе)
}
