package main

import "fmt"

//11) Дано число n. Определите, каким по счету числом Фибоначчи оно является, Если n не является числом Фибоначчи, выведите число 0.
func main() {
	var number int
	fmt.Println("Введите число для определения, является ли оно числом Фибоначчи")
	fmt.Scan(&number)
	prev, current, i := 0, 1, 1
	for ; current < number; i++ {
		prev, current = current, current+prev
	}
	if current == number {
		fmt.Println(i)
	} else {
		fmt.Println(0)
	}
}
