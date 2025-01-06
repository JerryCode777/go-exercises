package main

import "fmt"

func main() {
	var number, value int
	fmt.Print("Enter the number of values that will be in the series: ")
	fmt.Scan(&number)

	var max int
	for i := 1; i <= number; i++ {
		fmt.Print("Enter a value of the series: ")
		fmt.Scan(&value)
		if max < value {
			max = value
		}
	}
	fmt.Println("The maximum value of the series is: ", max)
}
