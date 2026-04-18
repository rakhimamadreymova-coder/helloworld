package main

import (
	"fmt"
)

func ExampleMaskCard() {
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Panic:", r)
			}
		}()
		fmt.Println(MaskCard("1234567890123456"))
	}()

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Panic:", r)
			}
		}()
		fmt.Println(MaskCard("1234-5678-9012-3456"))
	}()

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Panic:", r)
			}
		}()
		fmt.Println(MaskCard("123456789012345"))
	}()

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Panic:", r)
			}
		}()
		fmt.Println(MaskCard("12345678901234567"))
	}()

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Panic:", r)
			}
		}()
		fmt.Println(MaskCard("1234abcd90123456"))
	}()

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Panic:", r)
			}
		}()
		fmt.Println(MaskCard(""))
	}()

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Panic:", r)
			}
		}()
		fmt.Println(MaskCard("----"))
	}()

	// Output:
	// 1234 **** **** 3456
	// 1234 **** **** 3456
	// Panic: неверная длина: ожидается 16 цифр, получает 15
	// Panic: неверная длина: ожидается 16 цифр, получвет 17
	// Panic: неверный символ 'a' на позиции 5
	// Panic: неверная длина: ожидается 16 цифр, получает 0
	// Panic: неверная длина: ожидается 16 цифр, получает 0
}
