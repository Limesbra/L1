package main

import (
	"fmt"
	"task_8/internal/bit"
)

func main() {
	n := int64(1111111111)

	bit.SetBit(&n, 0, 0)
	fmt.Println(n)

	bit.SetBit(&n, 34, 1)
	fmt.Println(n)

	bit.SetBit(&n, 63, 1)
	fmt.Println(n)

	bit.SetBit(&n, 63, 0)
	fmt.Println(n)
}
