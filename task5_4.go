package main

import (
	"fmt"
	"math"
)

//4) Допустим у нас есть  осциллятор, представляющий собой механическую систему, состоящую из материальной точки на конце невесомой нерастяжимой нити или лёгкого стержня и находящуюся в однородном поле сил тяготения в вакууме при ну.  Вычислить  вычислить период колебаний. Результат должен быть представлен как структура с публичным методом

type Pendulum struct {
	Length float64
	G      float64
}

func (pendulum Pendulum) Oscillation() float64 {
	g := pendulum.G
	if g == 0 {
		return 0
	}
	return 2 * math.Pi * math.Sqrt(pendulum.Length/g)
}
func main() {
	fmt.Println("Введите длину маятника в метрах и ускорение свободного ускорения в м/с")
	pendulum := Pendulum{}
	fmt.Scan(&pendulum.Length, &pendulum.G)
	fmt.Printf("Период колебания равен %.2f с", pendulum.Oscillation())
}
