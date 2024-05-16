package main

import (
	"fmt"
	"unicode"
)

// проверка символов на уникальность, игнорирующая регистр
func IsUnique(s string) bool {
	runes := []rune(s)
	m := make(map[rune]bool)
	if len(runes) == 0 {
		return false
	}
	for _, r := range runes {
		// приводим все к нижнему регистру
		r = unicode.ToLower(r)
		if m[r] {
			return false
		}
		m[r] = true
	}
	return true

}

func main() {
	s1 := "abcd"           // true
	s2 := "abCdefAaf"      //false
	s3 := "😁😁"             // false
	s4 := "😁😂😃😄😅1234"      // true
	s5 := "  "             // false
	s6 := ""               // false
	s7 := "МАма МЫЛА Раму" // false

	fmt.Println(s1, "-", IsUnique(s1))
	fmt.Println(s2, "-", IsUnique(s2))
	fmt.Println(s3, "-", IsUnique(s3))
	fmt.Println(s4, "-", IsUnique(s4))
	fmt.Println(s5, "-", IsUnique(s5))
	fmt.Println(s6, "-", IsUnique(s6))
	fmt.Println(s7, "-", IsUnique(s7))

}
