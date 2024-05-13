package quicksort

// Comparable - интерфейс, представляющий типы, которые могут сравниваться и сортироваться.
type Comparable interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 | string
}

// QuickSort - быстрая сортировка
// Принимает массив и индекс опорного элемента
// Возвращает отсортированный массив
func QuickSort[T Comparable](arr []T, p int) []T {
	// проверяем что индекс опорного элемента не выходит за пределы массива
	// и что длина массива больше единицы
	if p < 0 || p >= len(arr) {
		return arr
	} else if len(arr) <= 1 {
		return arr
	}
	// инициализируем левую и правую границы
	l := 0
	r := len(arr) - 1
	// разделения массива по опорному элементу
	for l <= r {
		for arr[l] < arr[p] {
			l++
		}
		for arr[r] > arr[p] {
			r--
		}
		// если условие выполняется меняем местами элементы
		if l <= r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}

	// рекурсивно вызываем функцию для левой и правой части массива
	if r > 0 {
		QuickSort(arr[:r+1], r)
	}
	if l < len(arr)-1 {
		QuickSort(arr[l:], len(arr[l:])-1)
	}
	return arr
}
