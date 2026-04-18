package main

import "fmt"

func main() {
	var amount int64
	isAzo := true

	fmt.Print("Введите сумму перевода")
	fmt.Scan(&amount)

	var paidPart int64

	if isAzo {
		paidPart = amount - 1_000_000
		if paidPart < 0 {
			paidPart = 0 // если сумма меньше 1 000 000 — всё бесплатно
		}
	} else {
		paidPart = amount
	}

	commission := paidPart * 45 / 10000
	total := amount + commission

	fmt.Println("Сумма перевода:  ", amount, "сум")
	if isAzo {
		fmt.Println("Бесплатно (Аьзо): 1 000 000 сум")
		fmt.Println("Платная часть:   ", paidPart, "сум")
	}
	fmt.Println("Комиссия:        ", commission, "сум")
	fmt.Println("Итого спишется:  ", total, "сум")
}
