package main

//========= Чек =========
//Отправитель: Sherali
//Получатель: Alisher
//Сумма: 500000 сум
//Комиссия: 5000 сум
//Итого: 505000 сум
//=======================

import "fmt"

func main() {
	sender := "Sherali"
	receiver := "Alisher"
	amount := 500000

	commission := amount / 100
	total := amount + commission

	fmt.Printf("========== Чек ==========\n")
	fmt.Printf("Отправитель: %s\n", sender)
	fmt.Printf("Получатель: %s\n", receiver)
	fmt.Printf("Сумма: %d сум\n", amount)
	fmt.Printf("Комиссия: %d сум\n", commission)
	fmt.Printf("Итого: %d сум\n", total)
	fmt.Printf("========================\n")
}
