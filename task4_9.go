package main

import "fmt"

//9) Вводится одно число n. Вывести со склонением n машин
func main() {
	var number int
	fmt.Scan(&number)
	switch {
	case number%100/10 == 1:
		fmt.Printf("%d машин", number)
	case number%10 == 1:
		fmt.Printf("%d машина", number)
	case number%10 > 4:
		fmt.Printf("%d машин", number)
	default:
		fmt.Printf("%d машины", number)
	}
}
