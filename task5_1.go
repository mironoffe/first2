package main

import "fmt"

type Point struct {
	x, y int
}

//1) функция. передается N структур у которых есть x и y.   0> N < 5. Вычислить площадь фигуры в 2х мерном пространстве
func main() {
	var count int
	fmt.Println("Введите количество точек")
	fmt.Scan(&count)
	points := make([]Point, count)
	if count <= 2 {
		fmt.Printf("Площадь фигуры из %d точек всегда равна 0", count)
		return
	}
	fmt.Printf("Начинайте ввод координат точек в формате x1 y1 x2 y2 ... x%d y%d \n", count, count)
	for i := range points {
		fmt.Scan(&points[i].x, &points[i].y)
	}
	fmt.Println("Площадь фигуры равна ", calcSquare(points...))
}

func calcSquare(points ...Point) int {
	square := 0
	pointsCount := len(points)
	for i := range points {
		j := (i + 1) % pointsCount
		square += (points[j].x - points[i].x) * (points[j].y + points[i].y) / 2
	}
	if square < 0 {
		square *= -1
	}
	return square
}
