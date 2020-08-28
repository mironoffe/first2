package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//4)Мне нужна функция которая на вход получает целое положительное число (uint), возвращает число того же типа в котором отсутствуют нечетные цифры или цифра 0, если же получившееся число равно 0, то возвращается число 100. Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе
//Возвращяемый тип uint
func main() {
	var number uint = 765301186112340864
	var start time.Time

	start = time.Now()
	fmt.Println(replace(number))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println(replace2(number))
	fmt.Println(time.Since(start))
}

func replace(number uint) uint {
	var res uint = 0
	var multiple uint = 1
	for {
		digit := number % 10
		number = number / 10
		if digit%2 == 0 && digit != 0 {
			res += digit * multiple
			multiple *= 10
		}
		if number == 0 {
			break
		}
	}
	if res == 0 {
		return 100
	}
	return res
}

func replace2(number uint) uint {
	str := fmt.Sprint(number)
	res, _ := strconv.ParseUint(strings.Map(func(r rune) rune {
		if int(r)%2 == 1 || r == '0' {
			return -1
		}
		return r
	}, str), 10, 64)
	return uint(res)
}
