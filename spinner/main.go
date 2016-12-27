package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func worker(x int) int {
	// sleep for 3 seconds to imitate heavy workload
	time.Sleep(3 * time.Second)
	return x * x
}

func main() {
	go spinner(100 * time.Millisecond)
	fmt.Printf("\rResult: %d\n", worker(10))
}
