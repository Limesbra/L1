package main

import (
	"fmt"
	"task_11/internal/intersection"
)

func main() {
	num1 := []int{-1, 77, 8, 4, 55, 6, 7, 8, 9, -10, 55}
	num2 := []int{6, 7, 8, -9, 13, 45, 67, 67, 67, -1}

	fmt.Println(intersection.Intersection(num1, num2)) // пересечение [-1 6 7 8]
}
