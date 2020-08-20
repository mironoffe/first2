package main

import "fmt"

//2)  Подается число N;  и подется  последовательность X  длиной  N; сумму двузначных чисел в последовательности X , кратных 8; N - кол-во эллементов которое приходит
func main() {
	var (
		length int
		n      int
		sum    = 0
	)
	fmt.Println("Введите количество элементов в последовательности")
	fmt.Scan(&length)
	fmt.Println("Введите элементы")
	for i := 0; i < length; i++ {
		fmt.Scan(&n)
		if n < 100 && n > 9 && n%8 == 0 {
			sum += n
		}
	}
	fmt.Println("Сумма двузначных чисел кратных 8 равна", sum)
}
