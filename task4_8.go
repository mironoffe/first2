package main

import "fmt"

//8) Вводится одно  число n, Вывести цифровой корень числа n.
func main() {
	var number int
	fmt.Scan(&number)
	if number < 0 {
		number *= -1
	}
	for ; number > 9; number = digitSum(number) {

	}
	fmt.Println(number)
}

func digitSum(number int) int {
	sum := 0
	for {
		sum, number = sum+number%10, number/10
		if number == 0 {
			return sum
		}
	}
}
