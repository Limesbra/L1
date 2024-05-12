package binarysearch

import "golang.org/x/exp/constraints"

// BinarySearch - бинарный поиск
// Принимает массив и значение для поиска
// Возвращает индекс(искомого элемента или индекс на котором остановился бинарный поиск) и булевое значение(true - если элемент найден, false - если нет)
func BinarySearch[T constraints.Ordered](arr []T, target T) (uint, bool) {
	// Левая и правая границы
	var (
		left  uint = 0
		right uint = uint(len(arr) - 1)
	)
	//Обрабатываем случаи, когда массив пустой или с одним элементом
	if right == 0 {
		return right, false
	} else if right == 1 {
		if arr[0] == target {
			return 0, true
		} else {
			return right, false
		}
	}
	// бинарный поиск
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return uint(mid), true
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right, false
}
