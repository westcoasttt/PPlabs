package main

import "fmt"

func main() {
  var num1, num2 int

  fmt.Print("Введите первое число: ")
  fmt.Scanln(&num1)

  fmt.Print("Введите второе число: ")
  fmt.Scanln(&num2)

  // Сложение
  sum := num1 + num2
  fmt.Println("Сумма:", sum)

  // Вычитание
  difference := num1 - num2
  fmt.Println("Разность:", difference)

  // Умножение
  product := num1 * num2
  fmt.Println("Произведение:", product)

  // Деление
  if num2 != 0 {
    quotient := num1 / num2
    fmt.Println("Частное:", quotient)
  } else {
    fmt.Println("Деление на ноль невозможно!")
  }

  // Остаток от деления
  remainder := num1 % num2
  fmt.Println("Остаток от деления:", remainder)
}
