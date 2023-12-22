package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib := -1
	var mem = map[int]int{
		0: 0,
		1: 1,
	}
	
	return func() int {
		fib += 1
		
		if fib == 0 {
			return 0
		}
		if fib == 1 {
			return 1
		}
		
		newVal := mem[fib - 1] + mem[fib - 2]
		mem[fib] = newVal
		
		return newVal
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
