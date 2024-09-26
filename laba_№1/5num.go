package main

import "fmt"

// Функция для вычисления суммы двух чисел с плавающей запятой
func sum(num1, num2 float64) float64 {
 return num1 + num2
}

// Функция для вычисления разности двух чисел с плавающей запятой
func difference(num1, num2 float64) float64 {
 return num1 - num2
}

func main() {
 var num1, num2 float64

 fmt.Print("Введите первое число: ")
 fmt.Scanln(&num1)

 fmt.Print("Введите второе число: ")
 fmt.Scanln(&num2)

 // Вывод результатов
 fmt.Println("Сумма:", sum(num1, num2))
 fmt.Println("Разность:", difference(num1, num2))
}
