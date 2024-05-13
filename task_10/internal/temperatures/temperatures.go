package temperatures

// Group - функция группирует температуру по диапазонам
// t - массив температур
// возвращает map[int][]float64
func Group(t []float64) map[int][]float64 {
	// создаем мапу
	temp := make(map[int][]float64)
	// перебираем массив
	for _, item := range t {
		// определяем ключ
		key := int(item/10) * 10
		// добавляем в мапу
		temp[key] = append(temp[key], item)
	}
	return temp
}
