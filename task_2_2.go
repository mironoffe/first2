package main

import "fmt"

//2 Считать 2 переменные и сложить их
func main() {
	var a, b int
	fmt.Println("Enter an integer number:")
	fmt.Scan(&a, &b)
	fmt.Println("The sum is", a*b)
}
