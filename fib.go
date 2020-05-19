package fib

import "fmt"

func Fibonacci() func() int {
	x, y := 0, 1
	return func() (r int) {
		r = x
		x, y = y, x + y
		return
	}
}

func Run() {
	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}