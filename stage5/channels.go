package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTask(channel chan int) {
	delay := rand.Intn(5)
	fmt.Println("Starting long task")
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Println("Long task finished")
	// write to the channel
	channel <- delay
}

func main() {
	rand.Seed(time.Now().Unix())
	// create a channel
	channel := make(chan int)
	for i := 1; i <= 3; i++ {
		go longTask(channel)
	}
	// read from the channel
	for i := 1; i <= 3; i++ {
		fmt.Println("Took", <-channel, "seconds")
	}
}
