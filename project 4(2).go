//===== Alifshop =====
//Товар:    iPhone 15 Pro
//Бренд:    Apple
//Цена:     12 990 000 сум
//В наличии: true
//Рассрочка: 12 мес → 1 082 500 сум/мес
//====================

package main

import "fmt"

func main() {
	name := "iPhone 15 Pro"
	brand := "Apple"
	var price int64 = 12990000
	inStock := true

	monthly := price / 12
	fmt.Printf("===== Alifshop =====\n")
	fmt.Printf("товар: %s\n", name)
	fmt.Printf("brand: %s\n", brand)
	fmt.Printf("цена: %d сум\n", price)
	fmt.Printf("В наличии: %t\n", inStock)
	fmt.Printf("Рассрочка: 12 мес → %d сум/мес\n", monthly)
	fmt.Printf("====================\n")

}
