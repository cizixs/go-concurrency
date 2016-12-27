# Spinner

This example demostrates how to run two goroutines for different tasks at the same time.

Main goroutine creates a child goroutine to display spinner(naive progress bar),
then the main goroutine does some heavy computing(simple sleep a certain time).

NOTE: This example is from *The Go Programming Lauguage* book.
