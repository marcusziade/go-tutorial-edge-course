package main

import "math/rand"

func CalculateValue(values chan int) {
	value := rand.Intn(10)
	println("Value calculated: ", value)
	values <- value
}

func main() {
	values := make(chan int)
	go CalculateValue(values)

	value := <- values
	println(value)
}