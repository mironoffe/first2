package main

import "fmt"

//14) Дан массив  из N целых чисел. Выведите +, если массив является красивой горой, иначе  -.
func main() {
	fmt.Println("Введите 10 чисел")
	var arr [10]int
	for i, _ := range arr {
		fmt.Scan(&arr[i])
	}

	climbing := true
	previous := arr[0]
	for _, v := range arr {
		if v < previous && climbing {
			climbing = false
		}
		if v > previous && !climbing {
			fmt.Println("-")
			return
		}
		previous = v
	}
	if climbing {
		fmt.Println("-")
	} else {
		fmt.Println("+")
	}
}
