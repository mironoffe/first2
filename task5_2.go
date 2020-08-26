package main

import (
	"fmt"
)

//2) Считать  строку.  и вставить после каждого символа [a-z]   (нижнее подчерк)_
func main() {
	var str string
	fmt.Scan(&str)

	//for i := int('a'); i <= int('z'); i++ {
	//	str = strings.ReplaceAll(str, string(i), string(i)+"_")
	//}
	//
	result := ""
	for _, symbol := range str {
		result += string(symbol)
		if symbol >= 'a' && symbol <= 'z' {
			result += "_"
		}
	}
	fmt.Println(result)
}
