package main

import (
	"fmt"
	"math"
)

//6 На ввод подается число секунд,  нужно вывести сколько это часов, минут и секунд
func main() {
	var a int
	fmt.Scan(&a)
	seconds := a % 60
	minutes := int(math.Floor(float64(a)/60)) % 60
	hours := int(math.Floor(float64(a)/60/60)) % 60
	fmt.Println(hours, minutes, seconds)
}
