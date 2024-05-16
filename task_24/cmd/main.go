package main

import (
	"fmt"
	"task_24/internal/point"
)

func main() {
	p1 := point.NewPoint(0.5, 6)
	p2 := point.NewPoint(-7, 18)
	fmt.Println(point.Distance(p1, p2)) // 14.150971698084906
}
