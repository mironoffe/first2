package main

import "fmt"

//5) Дано трехзначное число. Переверните его, а затем выведите.
func main() {
	var number int
	fmt.Scan(&number)
	invert := 0
	for {
		invert = invert*10 + number%10
		number = number / 10
		if number == 0 {
			break
		}
	}
	fmt.Println(invert)
}
