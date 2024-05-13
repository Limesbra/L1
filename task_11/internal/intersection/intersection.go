package intersection

// Comparable - интерфейс, представляющий типы, которые могут сравниваться и сортироваться.
type Comparable interface {
	int | int8 | int16 | int32 | int64 | float32 |
		float64 | string | uint | uint8 | uint16 |
		uint32 | uint64
}

// Intersection - пересечение двух множеств
// Принимает два множества и возвращает их пересечение
// Если пересечения нет, то возвращается пустой слайс
func Intersection[T Comparable](nums1 []T, nums2 []T) []T {
	set := make(map[T]int)
	// Заполняем мапу данными из первого множества
	for _, num := range nums1 {
		set[num] = 1
	}
	// Увеличиваем счетчик, если число есть во втором множестве
	for _, num := range nums2 {
		if set[num] != 0 && set[num] > 0 {
			set[num] += 1
		} else { // Защита от повторяющихся чисел во втором монжестве
			set[num] = -1
		}

	}

	result := make([]T, 0)

	// Заполняем слайс значениями которые есть в обоих множествах
	for key, value := range set {
		if value >= 2 {
			result = append(result, key)
		}
	}

	return result
}
