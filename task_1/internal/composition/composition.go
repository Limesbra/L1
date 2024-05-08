package composition

import (
	"task_1/internal/human"
)

type Action struct {
	human.Human // наследование через композицию
}
