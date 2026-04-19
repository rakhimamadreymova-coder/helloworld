package main

import (
	"fmt"
	"project5-2/internal/commission"
	"strings"
)

func main() {
	var senderFirst, senderLast string
	var receiverFirst, receiverLast string
	var cardNumber string
	var amount int
	var cardType int

	fmt.Print("Имя отправителя: ")
	fmt.Scan(&senderFirst)
	fmt.Print("Фамилия отправителя: ")
	fmt.Scan(&senderLast)
	fmt.Print("Имя получателя: ")
	fmt.Scan(&receiverFirst)
	fmt.Print("Фамилия получателя: ")
	fmt.Scan(&receiverLast)
	fmt.Print("Номер карты получателя: ")
	fmt.Scan(&cardNumber)
	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)

	if amount < 500 || amount > 15_000_000 {
		fmt.Println("Ошибка: сумма должна быть от 500 до 15 000 000 сум")
		return
	}

	fmt.Print("Alif карта? (1-да/0-нет): ")
	fmt.Scan(&cardType)

	isAlif := cardType == 1
	comm := commission.Calculate(amount, isAlif)
	total := amount + comm
	masked := commission.MaskCard(cardNumber)
	txn := commission.GenerateTransaction()
	dt := commission.GetDateTime()

	fmt.Println()
	fmt.Println("========= Электронный чек =========")
	fmt.Printf("Отправитель:       %s %s\n", strings.ToUpper(senderLast), strings.ToUpper(senderFirst))
	fmt.Printf("Получатель:        %s %s\n", strings.ToUpper(receiverLast), strings.ToUpper(receiverFirst))
	fmt.Printf("Номер транзакции:  %d\n", txn)
	fmt.Printf("Счёт зачисления:   %s\n", masked)
	fmt.Printf("Дата и время:      %s\n", dt)
	fmt.Printf("Сумма:             %d сум\n", amount)
	fmt.Printf("Комиссия:          %d сум\n", comm)
	fmt.Printf("Итого:             %d сум\n", total)
	fmt.Println("АО ALIF TECH · Лицензия ЦБ РУз № 000010")
	fmt.Printf("Статус:            Исполнено\n")
	fmt.Println("==================================")
}
