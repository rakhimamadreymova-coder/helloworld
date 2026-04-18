package main

import "fmt"

func main() {
	sender := "Sherali"
	receiver := "Alisher"
	amount := int64(500_000)
	isAzo := true

	fmt.Printf("Отправитель: %s\n", sender)
	fmt.Printf("Получатель: %s\n", receiver)
	fmt.Printf("сумма: %d\n", amount)
	fmt.Printf("Статус Аьзо: %t\n", isAzo)

}
