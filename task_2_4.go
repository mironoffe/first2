package main

import "fmt"

//4 Считать с клавиатуры число и вывести последнюю ее цифру
func main() {
	var a string
	fmt.Println("Enter an integer number:")
	fmt.Scan(&a)
	fmt.Println("The last digit is", a[len(a)-1:])
}
