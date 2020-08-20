package main

import "fmt"

//4) Подается числа A, B, С;  Нвдо  вывести  первое число кратное С из промежутка A - B
func main() {
	var (
		start     int
		end       int
		delimiter int
	)
	fmt.Println("Введите начало и конец числового промежутка и делитель")
	fmt.Scan(&start, &end, &delimiter)
	if start > end {
		tmp := start
		start = end
		end = tmp
	}
	if delimiter == 0 {
		fmt.Println("Делитель не может быть нулём")
	} else {
		result := ((start-1)/delimiter + 1) * delimiter
		if result > end {
			fmt.Println("В промежутке", start, "и", end, "нет числа кратного", delimiter)
		} else {
			fmt.Println("В промежутке", start, "и", end, "первое число кратное", delimiter, "является", result)
		}
	}
}
