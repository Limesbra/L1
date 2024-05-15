package animal

import "fmt"

type Fish struct {
}

type Bird struct {
}

type Cat struct {
}

func (f *Fish) Status() {
	fmt.Println("Агрессивно смотрит")
}

func (b *Bird) Status() {
	fmt.Println("Поет песню")
}

func (c *Cat) Status() {
	fmt.Println("Спит")
}
