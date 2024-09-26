package main

import "fmt"

func main() {
	var num1, num2, num3 float64

	fmt.Print("Введите первое число:")
	fmt.Scanln(&num1)

	fmt.Print("Введите второе число:")
	fmt.Scanln(&num2)

	fmt.Print("Введите третье число:")
	fmt.Scanln(&num3)

	result := (num1 + num2 + num3) / 3
	fmt.Println("Среднее значение 3-х чисел: ", result)


}