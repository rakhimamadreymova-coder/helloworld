package commission

import (
	"fmt"
	"math/rand"
	"time"
)

func Calculate(amount int, isAlif bool) int {
	if isAlif {
		return 0
	}
	return int(float64(amount) * 0.0029)
}

func MaskCard(cardNumber string) string {
	return "UZCARD**" + cardNumber[len(cardNumber)-4:]
}

func GenerateTransaction() int {
	return rand.Intn(900000000) + 100000000
}

func GetDateTime() string {
	return time.Now().Format("02.01.2006 15:04")
}

func FormatAmount(amount int) string {
	return fmt.Sprintf("%d %03d", amount/1000, amount%1000)
}
