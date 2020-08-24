package main

import "fmt"

//6) Заданы три числа - a, b, c  - длины сторон треугольника. Нужно проверить, является ли треугольник прямоугольным треугольником. Если является, вывести "Да". Иначе вывести "Нет"
func main() {
	var a, b, c int
	fmt.Println("Введите длины сторон треугольника")
	fmt.Scan(&a, &b, &c)

	if c < a {
		c, a = a, c
	}
	if c < b {
		c, b = b, c
	}
	if a <= 0 || b <= 0 || c <= 0 {
		fmt.Printf("Нет")
	} else if c*c == (a*a + b*b) {
		fmt.Println("Да")
	} else {
		fmt.Println("Нет")
	}
}
