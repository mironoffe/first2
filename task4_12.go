package main

import "fmt"

//12) Дано число N. Выведите его представление в двоичном и шестнацатеричном виде.
func main() {
	var number int
	fmt.Scan(&number)
	fmt.Printf("%b, %x", number, number)
}
