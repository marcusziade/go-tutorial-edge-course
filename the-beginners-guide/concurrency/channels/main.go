package main

import (
	"math/rand"
	"time"
)

func CalculateValue(values chan int) {
	for i := 0; i <= 10; i++ {
		value := rand.Intn(10)
		println("Value calculated: ", value)
		values <- value
	}
}

func main() {
	values := make(chan int, 2)
	go CalculateValue(values)

	for i := 0; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		value := <-values
		println(value)
	}
}
