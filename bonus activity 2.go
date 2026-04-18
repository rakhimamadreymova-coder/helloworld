package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	sender := flag.String("sender", "", "Имя отправителя")
	amount := flag.Int("amount", 0, "Сумма")
	flag.Parse()

	name := strings.ToUpper(*sender)

	thousands := *amount / 1000
	rest := *amount % 1000

	amountStr := ""
	if thousands > 0 {
		amountStr = fmt.Sprintf("%d %03d", thousands, rest)
	} else {
		amountStr = fmt.Sprintf("%d", *amount)
	}

	fmt.Println(amountStr + " сум поступили юхуу")
	fmt.Println(name + " отправил вам деньги")
}
