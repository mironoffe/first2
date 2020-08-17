package main

import "fmt"

//4 Считать с клавиатуры число и вывести последнюю ее цифру
func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(a % 10)
}
