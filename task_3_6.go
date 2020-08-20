package main

import "fmt"

//6)  Подается числа A, B, C; A - сумма кредита; B - процент; C - срок (в месяцах);  процент на процент;  Нужно вывести итоговую сумму которую нужно будет отдать
func main() {
	var (
		sum        int
		percent    float64
		monthCount int
	)
	fmt.Println("Введите сумму кредита в рублях, годовой процент и срок в месяцах")
	fmt.Scan(&sum, &percent, &monthCount)
	monthlyFactor := percent / 12 / 100
	totalFactor := 1.0
	for month := 0; month < monthCount; month++ {
		totalFactor *= 1 + monthlyFactor
	}
	coefficient := monthlyFactor * totalFactor / (totalFactor - 1)
	monthlyPay := coefficient * float64(sum)
	total := monthlyPay * float64(monthCount)
	fmt.Println("Итоговая сумма платежа равна", int(total), "рублей", int(total*100)%100, "копеек")
}
