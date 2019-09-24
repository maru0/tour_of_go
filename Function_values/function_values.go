package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fn, fn1 := 0, 1
	return func() int {
		// 0,1,1,2,3 --
		fn, fn1 = fn1, fn+fn1
		return fn
	}
}

func main() {
	f := fibonacci()
	fmt.Println(0)
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
