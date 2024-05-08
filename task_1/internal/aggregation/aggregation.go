package aggregation

import (
	"task_1/internal/human"
)

type Action_ struct {
	Human human.Human // наследование через агрегацию
}

func (a *Action_) SayHello() {
	a.Human.SayHello()
}
