package main

import (
	"fmt"
	"sync"
	"time"
)

func wgSay(text string, wg *sync.WaitGroup) {

	// We access the wg Goroutine and we exec Done() when finishing with defer
	defer wg.Done()

	fmt.Println(text)
}

func channelSay(text string, channel chan string) {
	channel <- text
}

func main() {
	// One way for managing the Goroutines is with WaitGroup
	// sync pkg interacts with Goroutines in a primitive way, so is efficient
	var wg sync.WaitGroup // Most efficient way, but most complex for mantaining
	// wg is the var that sotres the Goroutines and executes them

	fmt.Println("Hello")

	// Here we add the Goroutine to the WaitGroup
	wg.Add(1)

	// go is the reserved keyword for executing the Goroutine concurrently
	go wgSay("World", &wg)

	// We wait until it is done. This says to the main Goroutine (the one which exec the whole code)
	// that it has to wait another Goroutines before ending
	wg.Wait()

	// The Goroutines often are written in lambda funcs
	go func(text string) {
		fmt.Println(text)
	}("Bye")

	// another way to make a wait on the Goroutines is with a sleep (but inefficient)
	time.Sleep(time.Second * 1) // Bad practice

	// If we don't set a sleeper, WaitGroup or channel, the Goroutine won't execute

	/* Channels: Less efficient but easier implementation */
	channel := make(chan string, 1)

	fmt.Println("Hello (bef channel)")

	go channelSay("Bye Channel", channel)

	fmt.Println(<-channel)
}
