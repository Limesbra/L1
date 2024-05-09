package square

func Square(n int, ch chan<- int) {
	ch <- n * n
}
