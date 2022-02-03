package main

import (
	"fmt"
	"time"
)

func main() {
	go printCalculating()
	for i := 0; i < 10000000000; i++ {
	} // dummy loop. mainスレッドで実行される.
}

func printCalculating() {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\rCalculating...%c", r) // \rがあるので左寄せになる.
			time.Sleep(100 * time.Millisecond)
		}
	}
}
