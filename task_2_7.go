package main

import "fmt"

//7 На ввод подается кол-во градусов оборота часовой стрелки с начала суток. Определи  сколько сейчас целых часов и минут.

func main() {
	var a int
	fmt.Scan(&a)
	hours := a % 360 / (360 / 12)
	minutes := a % 360 % (360 / 12) * 2
	fmt.Println(hours, minutes)

}
