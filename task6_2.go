package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
2) Написать  парсер строки который вынимает все числа мкладывает их и возвращяет результат
strconv
*/

func main() {
	//14+3+3+25+354+321+4
	str := "14jlk3lfsaklj3kl25jlk354lk321j4lknjknklfsda276"
	var start time.Time

	start = time.Now()
	fmt.Println(lineSum(str))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println(lineSum2(str))
	fmt.Println(time.Since(start))
}

func lineSum(str string) int {
	accumulator := 0
	sum := 0
	for _, r := range str {
		if r < '0' || r > '9' {
			if accumulator != 0 {
				sum, accumulator = accumulator+sum, 0
			}
			continue
		}
		accumulator = accumulator*10 + int(r) - 48
	}
	if accumulator != 0 {
		sum += accumulator
	}
	return sum
}
func lineSum2(str string) int {
	accumulator := ""
	sum := 0
	for _, r := range str {
		if r < '0' || r > '9' {
			if accumulator != "" {
				number, _ := strconv.Atoi(accumulator)
				sum, accumulator = sum+number, ""
			}
			continue
		}
		accumulator += string(r)
	}
	if accumulator != "" {
		number, _ := strconv.Atoi(accumulator)
		sum += number
	}
	return sum
}
