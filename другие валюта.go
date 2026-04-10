package main

import "fmt"

func main() {
	fmt.Println("выбери валюту:")
	fmt.Println("1-EUR")
	fmt.Println("2-$")
	fmt.Println("3-РУБЛЬ")

	var choice int
	fmt.Scan(&choice)

	var amount float64
	fmt.Println("введите сумму:")
	fmt.Scan(&amount)

	var result float64
	switch choice {
	case 1:
		result = amount * 14079.98
	case 2:
		result = amount * 12142.10
	case 3:
		result = amount * 150.55
	}

	fmt.Printf("%.2f = %.2f UZS\n", amount, result)
}
