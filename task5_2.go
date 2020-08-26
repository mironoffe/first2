package main

import "fmt"

//2) Считать  строку.  и вставить после каждого символа [a-z]   (нижнее подчерк)_
func main() {
	var str string
	fmt.Scan(&str)
	result := make([]rune, 0)
	for _, symbol := range str {
		result = append(result, symbol)
		if symbol >= 'a' && symbol <= 'z' {
			result = append(result, '_')
		}
	}
	fmt.Println(string(result))
}
