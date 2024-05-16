package point

import "math"

// структура точки
type Point struct {
	x float64
	y float64
}

// Функция конструктор
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Функция определения расстояния между двумя точками
func Distance(p1, p2 *Point) float64 {
	// AB = √(xb - xa)2 + (yb - ya)2
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))

}
