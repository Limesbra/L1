package main

import (
	"task_25/internal/sleep"
	"time"
)

// Реализовать собственную функцию sleep.

func main() {
	sleep.Sleep(3 * time.Second)
	sleep.Sleep2(3 * time.Second)
}
