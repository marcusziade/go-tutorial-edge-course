package main

import "time"

func HelloWorld(name string) {
	time.Sleep(1 * time.Second)
	println("Hello, ", name)
}

func main() {
	go HelloWorld("Marcus")
	println("I should print first")
	time.Sleep(2 * time.Second)
}
