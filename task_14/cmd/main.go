package main

import (
	"task_14/internal/method1"
	"task_14/internal/method2"
)

func main() {
	var x interface{} = "hello"
	method1.Method1(x)
	method2.Method2(x)

	x = 42
	method1.Method1(x)
	method2.Method2(x)

	x = make(chan bool)
	method1.Method1(x)
	method2.Method2(x)

	x = true

	method1.Method1(x)
	method2.Method2(x)
}
