package main

import "fmt"

//5 Подается число N; N > 999 && N < 10000; Нужно выыести число десятков
func main() {
	var a int
	fmt.Println("Enter an integer number (999 < N < 10000):")
	fmt.Scan(&a)
	fmt.Println("The penultimate number is", a/10%10)
}
