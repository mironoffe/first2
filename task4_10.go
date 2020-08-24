package main

import "fmt"

//10) Вводится одно число n. n < 10000 Вывести его буквами на русском языке
func main() {
	var number int
	fmt.Scan(&number)
	parts := [5]string{thousands(number), hundreds(number), dozens(number), teens(number), ones(number)}
	for _, part := range parts {
		if part != "" {
			fmt.Print(part + " ")
		}
	}
}

func ones(number int) string {
	digit := number % 10
	if number/10%10 == 1 {
		return ""
	}
	// один, два, три, четыре, пять, шесть, семь, восемь, девять
	switch digit {
	case 0:
		return ""
	case 1:
		return "один"
	case 2:
		return "два"
	case 3:
		return "три"
	case 4:
		return "четыре"
	case 5:
		return "пять"
	case 6:
		return "шесть"
	case 7:
		return "семь"
	case 8:
		return "восемь"
	case 9:
		return "девять"
	}
	return ""
}

func dozens(number int) string {
	digit := number / 10 % 10
	// двадцать, тридцать, сорок, ...N-десят, девяносто

	switch digit {
	case 0:
		return ""
	case 1:
		return ""
	case 2:
		return "двадцать"
	case 3:
		return "тридцать"
	case 4:
		return "сорок"
	case 9:
		return "девяносто"
	default:
		return ones(digit) + "десят"
	}
}

func teens(number int) string {
	// десять, одиннадцать, двенадцать, N-надцать, четырнадцать, N-надцать,
	if number/10%10 != 1 {
		return ""
	}
	digit := number % 10
	switch digit {
	case 0:
		return "десять"
	case 1:
		return "одиннадцать"
	case 2:
		return "двенадцать"
	case 4:
		return "четырнадцать"
	default:
		return ones(digit) + "надцать"
	}
}

func hundreds(number int) string {
	// сто, двести, триста, четыреста, ...N-сот,

	digit := number / 100 % 10
	switch digit {
	case 0:
		return ""
	case 1:
		return "сто"
	case 2:
		return "двести"
	case 3:
		return "триста"
	case 4:
		return "четыреста"
	default:
		return ones(digit) + "сот"
	}
}

func thousands(number int) string {
	// одна тысяча, две тысячи, три тысячи, четыре тысячи, ... N тысяч
	digit := number / 1000 % 10
	switch digit {
	case 0:
		return ""
	case 1:
		return "одна тысяча"
	case 2:
		return "две тысячи"
	case 3:
		return "три тысячи"
	case 4:
		return "четыре тысячи"
	default:
		return ones(digit) + " тысяч"
	}
}
