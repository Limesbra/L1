package counter

import "sync"

// структура - счетчик
type Counter struct {
	count int
	sync.Mutex
}

// Inc - увеличивает счетчик
func (c *Counter) Inc() {
	c.Lock()
	c.count++
	c.Unlock()
}

// Value - возвращает значение счетчика
func (c *Counter) Value() int {
	return c.count
}

