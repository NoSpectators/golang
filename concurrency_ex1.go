// a go routine is similar to a thread in other languages

package main

import (
	"fmt"
	"time"
)

// Spinner displays an animation while computing the 45th Fibonacci number.
func spinner(delay time.Duration) {
	for { //need 2 loops to keep iterating through chars
		for _, char := range `-\|/`{ //iterates through chars between backticks once
			fmt.Printf("\r%c", char)// `\r' keeps position of char, %c = unicode character
			time.Sleep(delay)		
		}
	}
	
}

//terribly inefficient recursive algorithm
func fib (x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func main() {
	go spinner(100*time.Millisecond) //rest of program doesn't wait for spinner to return
	const n = 45
	fibN := fib(n) //slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

