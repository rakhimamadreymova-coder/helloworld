package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	card := flag.String("card", "", "Номер карты")
	flag.Parse()

	normalized := strings.ReplaceAll(*card, " ", "")
	normalized = strings.ReplaceAll(normalized, "-", "")

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
