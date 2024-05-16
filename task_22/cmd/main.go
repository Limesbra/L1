package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	c := big.NewInt(1)

	// создаем два больших числа
	a := new(big.Int).SetUint64(math.MaxInt64 + 2)
	b := new(big.Int).SetUint64(math.MaxInt64 + 1)

	// сложение двух больших чисел
	res_add := new(big.Int).Add(a, b) // 18446744073709551617
	fmt.Println(res_add)

	// вычитание малого числа из большого
	res_sub := new(big.Int).Sub(a, c)
	fmt.Println(res_sub) // 9223372036854775808

	// вычитание большого числа из большого
	res_sub_another := new(big.Int).Sub(b, a)
	fmt.Println(res_sub_another) // -1

	// d := 9223372036854775808 // переполнение
	// fmt.Println(c)

	// f := 18446744073709551617 // переполнение
	// fmt.Println(c)
}
