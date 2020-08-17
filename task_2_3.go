package main

import "fmt"

//3 Считать переменую и возвести ее в вкадрат
func main() {
	var a int
	fmt.Println("Enter an integer number:")
	fmt.Scan(&a)
	fmt.Println("The square of number is", a*a)
}
