package main

import "fmt"

func main() {
	var amount int64

	fmt.Println("Введите сумму оплаты")
	fmt.Scan(&amount)

	cashbackServices := amount * 22 / 1000
	cashbackNasiya := amount * 1 / 1000

	fmt.Println("Кэшбек за сервисы(2.2%): ", cashbackServices)
	fmt.Println("Кэшбек за рассрочку(0.1%): ", cashbackNasiya)

}
