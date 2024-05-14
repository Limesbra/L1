package square

import (
	"sync"
)

// square - функция, которая возводит число в квадрат
func Square(n int, ch chan<- int, mu *sync.Mutex) {
	mu.Lock()
	square := n * n
	ch <- square
	mu.Unlock()
}
