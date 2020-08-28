package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// 3) Считать и распарсить csv строку содержащую 3 числа и вывести их сумму
//encoding/csv
func main() {
	reader := csv.NewReader(os.Stdin)
	reader.FieldsPerRecord = 3
	line, e := reader.Read()
	if e != nil {
		fmt.Println("Передана не CSV строка, либо в строке не 3 значения")
		return
	}
	result := 0.0
	for _, v := range line {
		number, e := strconv.ParseFloat(v, 64)
		if e != nil {
			fmt.Println("Передано не число:", v)
			return
		}
		result += number
	}
	fmt.Print("Сумма - ", result)
}
