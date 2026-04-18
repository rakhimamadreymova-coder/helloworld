package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string
	var amount int

	fmt.Print("Введите имя отправителя: ")
	fmt.Scanln(&name)

	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)

	name = strings.ToUpper(name)

	thousands := amount / 1000
	rest := amount % 1000

	amountStr := ""
	if thousands > 0 {
		amountStr = fmt.Sprintf("%d %03d", thousands, rest)
	} else {
		amountStr = fmt.Sprintf("%d", amount)
	}

	fmt.Println(amountStr + " сум поступили ураа")
	fmt.Println(name + " отправил вам деньги")
}
