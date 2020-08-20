package main

import "fmt"

//8) Подается числа A, B;  Нужно вывести квадрат суммы всех входящих чисел промежутка A-B; Скорость O(1)
func main() {
	var (
		start int
		end   int
	)
	fmt.Println("Введите начало и конец числового промежутка")
	fmt.Scan(&start, &end)
	sum := (end*(end+1) - start*(start-1)) / 2
	fmt.Println("Квадрат суммы всех входящих чисел промежутка", start, "и", end, "равна", sum*sum)
}
