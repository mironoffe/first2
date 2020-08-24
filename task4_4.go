package main

import "fmt"

//4) Дано трехзначное число. Найдите сумму его цифр.
func main() {
	var number int
	fmt.Scan(&number)
	sum := 0
	for {
		sum, number = sum+number%10, number/10
		if number == 0 {
			break
		}
	}
	fmt.Println(sum)
}
