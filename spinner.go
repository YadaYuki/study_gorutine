package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond) // 計算実行中に表示するspinner
	fmt.Printf("\rFibonacci(%d) = %d\n", 45, fib(45))
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
