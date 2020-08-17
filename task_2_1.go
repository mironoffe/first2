package main

import "fmt"

//1 Считать с клавиатуры *2 +1000
func main() {
	var a int
	fmt.Println("Enter an integer number:")
	fmt.Scan(&a)
	fmt.Println("The result on N * 2 + 1000 is", a*2+1000)
}
