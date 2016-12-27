package main

import (
	"fmt"
)

// sum accepts a slice and channel,
// adds all the numbers in slice, and send the result to channel
func sum(s []int, c chan int) {
	sum := 0
	for _, n := range s {
		sum += n
	}
	c <- sum
}

func fillSlice(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] = i + 1
	}
}

func main() {

	// create a slice of length 1000
	tasks := 10000
	s := make([]int, tasks)
	fillSlice(s)

	// how many go routines to run the task
	num := 10

	c := make(chan int)

	// how many tasks per worker
	// NOTE: please make sure tasks is evenly divided,
	// otherwise the result can be wrong
	perWorker := tasks / num

	// run the worker the do the hard working(calculating...)
	for i := 0; i < num; i++ {
		go sum(s[i*perWorker:(i+1)*perWorker], c)
	}

	// wait for results from workers
	sum := 0
	for i := 0; i < num; i++ {
		result := <-c
		sum += result
		fmt.Printf("Received %d\n", result)
	}

	fmt.Printf("1+2+...+%d = %d\n", tasks, sum)
}
