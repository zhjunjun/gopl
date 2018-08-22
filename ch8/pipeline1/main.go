// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 228.

// Pipeline1 demonstrates an infinite 3-stage pipeline.
package main

import (
	"fmt"
	"time"
)

//!+
func main() {
	naturals := make(chan int)
	squares := make(chan int)
	threef := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
			time.Sleep(1 * time.Second)
		}
	}()

	// Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// Threef
	go func(){
		for {
			x := <-squares
			threef <- x * x * x
		}
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}

//!-
