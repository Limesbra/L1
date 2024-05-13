package binarysearch

// Comparable - интерфейс, представляющий типы, которые могут сравниваться и сортироваться.
type Comparable interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 | string
}

// BinarySearch - бинарный поиск
// Принимает массив и значение для поиска
// Возвращает индекс(искомого элемента или индекс на котором остановился бинарный поиск) и булевое значение(true - если элемент найден, false - если нет)
func BinarySearch[T Comparable](arr []T, target T) (uint, bool) {
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
