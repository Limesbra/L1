package main

import (
	"fmt"
	"task_1/internal/aggregation"
	"task_1/internal/composition"
	"task_1/internal/human"
)

func main() {

	a := aggregation.Action_{
		Human: human.Human{
			Name:    "Denis",
			Surname: "Volkov",
			Age:     23,
		},
	}

	c := composition.Action{
		Human: human.Human{
			Name:    "Irina",
			Surname: "Novikova",
			Age:     87,
		},
	}

	a.SayHello()                                                                         // В агрегации поля и методы внутренней структуры недоступны через внешнюю структуру
	fmt.Println("Hello, my name is", c.Name, c.Surname, "\b. I am", c.Age, "years old.") // В композиции поля и методы внутренней структуры доступны через внешнюю структуру

}
