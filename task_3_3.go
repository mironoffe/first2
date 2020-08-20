package main

import "fmt"

//3) передается последовательность чисел X. Последовательность числе заканчивается 0. Количество элементов X последовательности, которые равны ее наибольшему элементу.
func main() {

	var (
		n        int
		max      int
		maxCount = 0
	)
	fmt.Println("Вводите любые числа. Для окончания введите 0")
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		if max < n {
			max = n
			maxCount = 0
		}
		if max == n {
			maxCount++
		}
	}
	rusCountName := "раз"
	if (maxCount+8)%10 < 3 && maxCount/10%10 != 1 {
		rusCountName = "раза"
	}
	fmt.Println("Наибольшее число встречалось", maxCount, rusCountName)
}
