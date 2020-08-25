package main

import "fmt"

//3) Считать  строку.   и опредилить полиндром ли она за 1 for
func main() {
	var str string
	fmt.Scan(&str)
	for i, j := 0, len(str)-1; i <= j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			fmt.Printf("Строка «%s» не полиндром", str)
			return
		}
	}
	fmt.Printf("Строка «%s» полиндром", str)
}
