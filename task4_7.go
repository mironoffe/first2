package main

import "fmt"

//7) Даны три числа a, b, c. Определите, существует ли треугольник с такими сторонами.
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
	} else if c < a+b {
		fmt.Println("Да")
	} else {
		fmt.Println("Нет")
	}
}
