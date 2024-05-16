package main

import (
	"fmt"
	"task_23/internal/remove"
)

func main() {

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{1, 2, 3, 4, 5}
	slice3 := []int{1, 2, 3, 4, 5}
	i := 2
	j := 3
	k := 1

	// удаляем i и j элементы из слайсов
	slice1 = remove.Remove1(slice1, i)
	slice2 = remove.Remove2(slice2, j)
	ns := remove.Remove3(slice3, k)

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3, ns)

}
