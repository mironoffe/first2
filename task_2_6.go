package main

import "fmt"

//6 На ввод подается число секунд,  нужно вывести сколько это часов, минут и секунд
func main() {
	var a int
	fmt.Scan(&a)
	seconds := a % 60
	minutes := a / 60 % 60
	hours := a / 60 / 60 % 60
	fmt.Println("Hours:", hours)
	fmt.Println("Minutes:", minutes)
	fmt.Println("Seconds:", seconds)
}
