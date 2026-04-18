package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Введите номер карты: ")

	var a, b, c, d string
	n, _ := fmt.Scan(&a, &b, &c, &d)

	input := a
	if n > 1 {
		input = a + b + c + d
	}

	normalized := strings.ReplaceAll(input, "-", "")

	if len(normalized) != 16 {
		fmt.Println("Ошибка: неверный формат карты")
		return
	}

	valid := true
	for i := 0; i < len(normalized); i++ {
		ch := normalized[i]
		if ch < '0' || ch > '9' {
			valid = false
			break
		}
	}

	if !valid {
		fmt.Println("Ошибка: неверный формат карты")
		return
	}

	fmt.Println("Нормализованный номер:", normalized)
	fmt.Println("Длина: 16")

	// Дополнительное задание
	first4 := ""
	last4 := ""

	for i := 0; i < len(normalized); i++ {
		if i < 4 {
			first4 = first4 + string(normalized[i])
		}
		if i >= 12 {
			last4 = last4 + string(normalized[i])
		}
	}

	masked := first4 + " **** **** " + last4
	fmt.Println("Карта:", masked)
}
