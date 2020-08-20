package main

import "fmt"

//1)  Подается число N;  Нужно вывести сумму арифметической прогресии числа N; Скорость O(1)
func main() {
	var a int
	fmt.Println("Введите число")
	fmt.Scan(&a)
	fmt.Println("Сумма арифметической прогрессии 1 ..", a, "равна", a*(a+1)/2)
}
