package main

import "fmt"

func main() {
	var num int

	fmt.Print("Введите целое число: ")
	fmt.Scan(&num)

	if num >= 12307 {
		fmt.Printf("Введенное число %d уже больше или равно 12307\n", num)
		return
	}

	fmt.Printf("Начальное число: %d\n", num)

	for num < 12307 {

		if num < 0 {
			num = num * -1

		} else if num%7 == 0 {
			num = num * 39

		} else if num%9 == 0 {
			num = num*13 + 1
			continue

		} else {
			num = (num + 2) * 3
		}

		if num%13 == 0 && num%9 == 0 {
			fmt.Print("service error")
		} else {
			num += 1
		}
	}

	fmt.Printf("Результат: %d", num)
}
