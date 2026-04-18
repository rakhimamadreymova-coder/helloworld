package main

import (
	"fmt"
	"strings"
)

func MaskCard(input string) string {
	normalized := strings.ReplaceAll(input, "-", "")

	if len(normalized) != 16 {
		panic(fmt.Sprintf("неверная длина: ожидается 16 цифр, получено %d", len(normalized)))
	}

	for i := 0; i < len(normalized); i++ {
		if normalized[i] < '0' || normalized[i] > '9' {
			panic(fmt.Sprintf("неверный символ '%c' на позиции %d", normalized[i], i+1))
		}
	}

	first4 := normalized[:4]
	last4 := normalized[12:]
	return first4 + " **** **** " + last4
}

func main() {
	fmt.Print("Введите номер карты: ")
	var a, b, c, d string
	n, _ := fmt.Scan(&a, &b, &c, &d)

	input := a
	if n > 1 {
		input = a + b + c + d
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ошибка:", r)
		}
	}()

	masked := MaskCard(input)
	fmt.Println("Карта:", masked)
}
