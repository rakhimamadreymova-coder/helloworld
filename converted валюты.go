package main

import "fmt"

func main() {
	var amount float64
	fmt.Print("Введите сумму в Euro: ")
	fmt.Scan(&amount)

	rate := 14079.98
	result := amount * rate

	fmt.Printf("%.2f EUR = %.2f UZS\n", amount, result)
}
