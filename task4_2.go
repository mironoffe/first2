package main

import (
	"fmt"
	"time"
)

//2) Дан массив из n целых чисел.
//Переместите все нули в конец, сохраняя относительный порядок ненулевых элементов. Вы должны сделать это в том же массиве, не создавая нового массива. Использовать удаление, добавение нельзя не более 1 for
func main() {
	var a = [6]int{1, 2, 3, 0, 0, 8}

	start := time.Now()

	lastPoint := 0
	for i, lenA := 0, len(a); i < lenA; i++ {
		if a[i] != 0 {
			if lastPoint != i {
				a[lastPoint], a[i] = a[i], a[lastPoint]
			}
			lastPoint++
		}
	}

	fmt.Println(a)
	fmt.Println(time.Since(start))
}
