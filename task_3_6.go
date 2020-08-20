package main

import "fmt"

//6)  Подается числа A, B, C; A - сумма кредита; B - процент; C - срок (в месяцах);  процент на процент;  Нужно вывести итоговую сумму которую нужно будет отдать
func main() {
	var (
		sum        int
		percent    int
		monthCount int
	)
	fmt.Println("Введите сумму кредита в рублях, годовой процент и срок в месяцах")
	fmt.Scan(&sum, &percent, &monthCount)
	lastYearMonthCount := monthCount % 12
	yearsCount := monthCount / 12
	for year := 0; year < yearsCount; year++ {
		sum = sum * (100 + percent) / 100
	}
	sum = sum * (100 + percent*lastYearMonthCount/12) / 100
	fmt.Println("Итоговая сумма платежа равна", sum, "руб.")
}
