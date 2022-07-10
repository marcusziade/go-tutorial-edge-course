package main

import (
	"fmt"
	"runtime"
)

func main() {
	createArray()
}

func createArray() {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	planets = append(planets, "Pluto")
	for _, planet := range planets {
		fmt.Println(planet)
	}
}

func createVariables() {
	myAge := 29
	myAge++
	myAge -= 15
	println(myAge)
}

func createIfStatement() {
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

func createConstants() {
	const whatif = "What if?"
	println(whatif)
}

func createSwitch() {
	customerHeight := 100
	customerAge := 10

	switch {
	case customerHeight >= 130 || customerAge >= 18:
		fmt.Println("Can access all rides")
	case customerHeight >= 120:
		fmt.Println("Can access children rides")
	default:
		fmt.Println("Can't access rides")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		println("macOS")
	case "linux":
		println("Linux machine")
	default:
		println("Something else")
	}
}
