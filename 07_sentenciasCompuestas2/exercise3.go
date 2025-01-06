package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter the number of numbers to input: ")
	fmt.Scan(&number)

	var value, greaterZero, lessZero, Zeros int

	for i := 0; i < number; i++ {
		fmt.Print("Enter a number: ")
		fmt.Scan(&value)
		if value > 0 {
			greaterZero++
		} else if value < 0 {
			lessZero++
		} else {
			Zeros++
		}
	}

	fmt.Println("There are ", greaterZero, " numbers greater than zero")
	fmt.Println("There are ", lessZero, " numbers less than zero")
	fmt.Println("There are ", Zeros, " zeros")

}
