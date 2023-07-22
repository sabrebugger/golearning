package main

import (
	"fmt"
)

func main() {
	var num1, num2 int

	fmt.Println("Введите первое число:")
	fmt.Scan(&num1)

	fmt.Println("Введите второе число:")
	fmt.Scan(&num2)

	sum := num1 + num2

	fmt.Printf("Сумма чисел %d и %d равна %d\n", num1, num2, sum)
}
