package remove

import "fmt"

// Remove1 - функция удаления элемента из слайса
// Если нам не важен порядок в слайсе
func Remove1(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		fmt.Println("Index is out of range")
		return nil
	}
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

// Remove2 - функция удаления элемента из слайса
// Если порядок элементов в слайсе важен
func Remove2(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		fmt.Println("Index is out of range")
		return nil
	}
	// объединяем два новых фрагмента вместе
	return append(slice[:i], slice[i+1:]...)
}

// Remove3 - функция удаления элемента из слайса
// Сохраняет исходный слайс
func Remove3(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		fmt.Println("Index is out of range")
		return nil
	}
	newSlice := make([]int, 0)
	newSlice = append(newSlice, slice[:i]...)
	return append(newSlice, slice[i+1:]...)
}
