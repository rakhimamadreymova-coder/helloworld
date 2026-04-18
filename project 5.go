package main

import (
	"fmt"
	"math"
)

//Введите сумму: 43925
//Alif карта? (1-да/2-нет): 0
//===== Чек =====
//услуга:    Madreymova Rakhima
//сумма:     43925 сум
//комиссия: 127 сум
//итого: 44052 сум
//Статус:: исполнено
//====================

func main() {
	var amount float64
	var alifChoice int

	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)

	if amount < 500 {
		fmt.Println("Ошибка: минимальная сумма 500 сум")
		return
	}
	if amount > 15_000_000 {
		fmt.Println("Ошибка: максимальная сумма 15 000 000 сум")
		return
	}

	fmt.Print("Alif карта? (1-да/0-нет): ")
	fmt.Scan(&alifChoice)

	isAlif := alifChoice == 1
	commission := calculateCommission(amount, isAlif)
	printReceipt("Madreymova Rakhima", amount, commission)
}

func calculateCommission(amount float64, isAlif bool) float64 {
	if isAlif {
		return 0
	}
	return math.Floor(amount * 0.0029)
}

func printReceipt(name string, amount float64, commission float64) {
	total := amount + commission
	fmt.Println("\n========= Чек =========")
	fmt.Printf("Услуга: %s\n", name)
	fmt.Printf("Сумма: %.0f сум\n", amount)
	fmt.Printf("Комиссия: %.0f сум\n", commission)
	fmt.Printf("Итого: %.0f сум\n", total)
	fmt.Println("Статус: Исполнено")
	fmt.Println("=======================")
}
