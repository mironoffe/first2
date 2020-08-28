package main

import (
	"fmt"
	"math"
	"math/rand"
)

/*
1) информация была сохранена в  map:

AllPrices := map[int][]string{
   10:   []string{...}, // товар  с ценой 10 -99
   100:  []string{...}, // товар  с ценой 100 - 999
   1000: []string{...}, // товар  с ценой более 1000
}

При запросе был получен map:
prices := map[string]int{
   товар: цена
}

Требуется чтобы map prices содержал только товары из группы  AllPrices[100]
*/

func main() {
	prices, AllPrices := genData(30)
	filteredPrices := make(map[string]int, len(AllPrices[100]))
	for _, name := range AllPrices[100] {
		if prices[name] != 0 {
			filteredPrices[name] = prices[name]
		}
	}
	fmt.Println(prices)
	fmt.Println(AllPrices)
	fmt.Println(filteredPrices)
}
func genData(pricesCount int) (prices map[string]int, AllPrices map[int][]string) {
	prices = make(map[string]int, pricesCount)
	for i := 0; i < pricesCount; i++ {
		for {
			name := GenName(rand.Intn(8) + 4)
			if prices[name] == 0 {
				prices[name] = int(math.Pow(2, float64(rand.Intn(14))))
				break
			}
		}
	}
	AllPrices = map[int][]string{
		10:   {},
		100:  {},
		1000: {},
	}
	for name, price := range prices {
		switch {
		case price >= 10000:
		case price >= 1000:
			AllPrices[1000] = append(AllPrices[1000], name)
		case price >= 100:
			AllPrices[100] = append(AllPrices[100], name)
		case price >= 10:
			AllPrices[10] = append(AllPrices[10], name)
		}
	}
	return prices, AllPrices
}
func GenName(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyz0123456789_"
	result := ""
	charsetLength := len(charset)
	for i := 0; i < length; i++ {
		result += string(charset[rand.Intn(charsetLength)])
	}
	return result
}
