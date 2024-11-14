package main

import (
	"fmt"
	"log"
)

// Целевой интерфейс
type PaymentProcessor interface {
	ProcessPayment(amount float64) (string, error)
}

// Реализация для Yandex Pay
type YandexPay struct{}

func (y *YandexPay) ProcessPayment(amount float64) (string, error) {
	if amount <= 0 {
		return "", fmt.Errorf("invalid amount for YandexPay: %.2f", amount)
	}
	return fmt.Sprintf("Processed $%.2f payment through YandexPay", amount), nil
}

// Реализация для SberPay
type SberPay struct{}

func (s *SberPay) ProcessPayment(amount float64) (string, error) {
	if amount <= 0 {
		return "", fmt.Errorf("invalid amount for SberPay: %.2f", amount)
	}
	return fmt.Sprintf("Processed $%.2f payment through SberPay", amount), nil
}

// Реализация для старой системы с другим интерфейсом
type OldMoneySystem struct{}

func (o *OldMoneySystem) MakePayment(amount float64) (string, error) {
	if amount <= 0 {
		return "", fmt.Errorf("invalid amount for Old Money System: %.2f", amount)
	}
	return fmt.Sprintf("Processed $%.2f payment through Old Payment System", amount), nil
}

// Адаптер для старой системы
type OldMoneyAdapter struct {
	oldSystem OldMoneySystem
}

func (a *OldMoneyAdapter) ProcessPayment(amount float64) (string, error) {
	return a.oldSystem.MakePayment(amount)
}

func main() {
	var processor PaymentProcessor

	// Используем Yandex Pay
	processor = &YandexPay{}
	if result, err := processor.ProcessPayment(100.00); err != nil {
		log.Printf("Error processing payment: %v", err)
	} else {
		fmt.Println(result)
	}

	// Используем SberPay
	processor = &SberPay{}
	if result, err := processor.ProcessPayment(200.00); err != nil {
		log.Printf("Error processing payment: %v", err)
	} else {
		fmt.Println(result)
	}

	// Используем для старой системы через адаптер
	oldSystem := OldMoneySystem{}
	processor = &OldMoneyAdapter{oldSystem: oldSystem}

	if result, err := processor.ProcessPayment(450.00); err != nil {
		log.Printf("Error processing payment: %v", err)
	} else {
		fmt.Println(result)
	}

	// Пример обработки ошибки с некорректной суммой
	if result, err := processor.ProcessPayment(-50.00); err != nil {
		log.Printf("Error processing payment: %v", err)
	} else {
		fmt.Println(result)
	}
}
