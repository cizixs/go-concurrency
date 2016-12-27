package main

import "fmt"

var battle = make(chan string)

func warrior(name string, done chan struct{}) {
	// select waits on multiple communication operations.
	// select will block until one of the cases can run, then it executes that case.
	// If multiple cases are ready, it chooses one at random.
	//
	// Since done is a blocked channel, each time only one goroutine can write to it.
	// So each goroutine tries to write message to it, if it succeeds(win), other goroutine
	// can read from it; otherwise someone has already written content to the channel(lose),
	// it can read from it.
	select {
	case opponent := <-battle:
		fmt.Printf("%s beat %s\n", name, opponent)
	case battle <- name:
	}
	done <- struct{}{}
}

func main() {
	// done is a channel to receive notifications from workers indicating that it has done processing.
	// struct{} is used because it does not cost memory space.
	// please refer to http://stackoverflow.com/questions/20793568/golang-anonymous-struct-and-empty-struct/41342407#41342407
	done := make(chan struct{})

	langs := []string{"Go", "C", "C++", "Java", "Perl", "Python"}
	for _, l := range langs {
		go warrior(l, done)
	}
	for _ = range langs {
		<-done
	}
}
