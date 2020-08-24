package main

import "fmt"

//3) На вход подается 5 чисел a, b, c, d, e
//x = a умножается на 2 b раз
//y =c умножается на 2 d раз
//сумма результата x и y делится на 2 e раз сделать сопоставимо по скорости
func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)
	fmt.Println((a<<b + c<<d) >> e)
}
