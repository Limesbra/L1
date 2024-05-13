package main

import (
	"fmt"
	"task_10/internal/temperatures"
)

func main() {
	t := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(temperatures.Group(t))
}
