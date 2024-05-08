package human

import "fmt"

type Human struct {
	Name    string
	Surname string
	Age     int
}

func (human *Human) SayHello() {
	fmt.Printf("Hello, my name is %s %s. I am %d years old.\n", human.Name, human.Surname, human.Age)
}
