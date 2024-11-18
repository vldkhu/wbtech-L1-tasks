package main

import (
	"fmt"
	"math/big"
)

func main() {

	a := new(big.Int)
	a.SetString("1234567890123456789", 10)
	b := new(big.Int)
	b.SetString("9876543210987654321", 10)

	//сложение
	sum := new(big.Int).Set(a)
	sum.Add(sum, b)
	fmt.Println("Сумма:", sum)

	//умножение
	multiplication := new(big.Int).Set(a) // копируем а в mul
	multiplication.Mul(multiplication, b)
	fmt.Println("Произведение:", multiplication)

	//деление
	division := new(big.Int).Set(b)
	division.Div(division, a)
	fmt.Println("Деление:", division)

	//вычитание
	subtraction := new(big.Int).Set(b)
	subtraction.Sub(subtraction, a)
	fmt.Println("Вычитание:", subtraction)
}
