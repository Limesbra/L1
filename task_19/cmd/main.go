package main

import (
	"fmt"
	"task_19/internal/reverse"
)

func main() {
	s := "The quick brown 狐 jumped over the lazy 犬"
	fmt.Println(reverse.Reverse(s))
	fmt.Println(reverse.Reverse2(s))
}
