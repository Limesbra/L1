package method2

import (
	"fmt"
	"reflect"
)

func Method2(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Println("Method2: i is of type", t)
}
