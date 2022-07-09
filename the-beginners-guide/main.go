package main

import (
	"fmt"
	"runtime"
)

func main() {
	// switch statements
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
