package main

import (
	"fmt"
	"math/big"
)

func sum(a, b *big.Int) *big.Int {
	res := new(big.Int) // выделяем память для результата
	res.Add(a, b)       // выполняем операцию
	return res          // возвращаем указатель на результат
}

func sub(a, b *big.Int) *big.Int {
	res := new(big.Int)
	res.Sub(a, b)
	return res
}

func multiply(a, b *big.Int) *big.Int {
	res := new(big.Int)
	res.Mul(a, b)
	return res
}

func division(a, b *big.Int) *big.Int {
	res := new(big.Int)
	res.Quo(a, b)
	return res
}

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("1234567890123456789", 10)
	b.SetString("9876543210987654321", 10)

	fmt.Printf("sum = %d\n", sum(a, b))
	fmt.Printf("subtraction = %d\n", sub(a, b))
	fmt.Printf("multiplication = %d\n", multiply(a, b))
	fmt.Printf("division = %d\n", division(a, b))
}
