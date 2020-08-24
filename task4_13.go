package main

import "fmt"

//13) На вход подается 2 числа a, b. Из числа a удалить заданную цифру b.
func main() {
	// todo
	var number, digit, result int
	fmt.Scan(&number, &digit)
	for i, j := 1, 1; i < number; i *= 10 {
		currentDigit := number / i % 10
		if currentDigit != digit {
			result += currentDigit * j
			j *= 10
		}
	}
	fmt.Println(result)
}
