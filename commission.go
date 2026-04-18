package main

/*commision*/

import (
	"fmt"
	"strings"
)

func maskPAN(input string) {
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
