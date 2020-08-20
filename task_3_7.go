package main

import "fmt"

//7) Подается числа A, B; Нужно вывести число которое  состоит из последовательного набора цифр входящих в оба числа, без повторения. В порядке нахождения в первом числе
func main() {
	var (
		result = ""
		first  = ""
		second = ""
	)
	fmt.Println("Введите 2 числа")
	fmt.Scan(&first, &second)
	for _, f := range first {
		alreadyExits := false
		for _, r := range result {
			if r == f {
				alreadyExits = true
				break
			}
		}
		if alreadyExits {
			continue
		}
		for _, s := range second {
			if s != f {
				continue
			}
			result += string(f)
			break
		}
	}
	if result == "" {
		fmt.Println("У введённых чисел нет общих цифр")
	} else {
		fmt.Println("Число которое состоит из последовательного набора цифр входящих в оба числа равно", result)
	}
}
