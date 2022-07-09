package main

import "fmt"

func main() {
	// if statements
	customerHeight := 130
	customerAge := 29

	if customerHeight >= 150 && customerAge >= 18 {
		fmt.Println("Can access")
	} else if customerHeight >= 120 {
		println("Can almost access")
	} else {
		fmt.Println("Can not access")
	}
}
