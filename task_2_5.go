package main

import (
	"fmt"
	"math"
)

//5 Подается число N; N > 999 && N < 10000; Нужно выыести число десятков
func main() {
	var a float64
	fmt.Scan(&a)
	fmt.Println(int(math.Floor(a/10)) % 10)
}
