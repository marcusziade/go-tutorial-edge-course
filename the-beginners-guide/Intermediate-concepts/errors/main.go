package main

import "errors"

var (
	ErrorIsLessThanZero = errors.New("Number is less than zero")
	ErrorIsNotEven      = errors.New("Number is not even xD")
)

func isEven(number int) error {
	if number%2 != 0 {
		return ErrorIsNotEven
	} else if number < 0 {
		return ErrorIsLessThanZero
	}

	return nil
}

func main() {
	error := isEven(-2)
	if error != nil {
		println(error.Error())
	}
}
