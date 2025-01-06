package main

import "fmt"

func main() {
	// 0:0:0 .... => .... 23.59.59

	for hours := 0; hours <= 23; hours++ {
		for minutes := 0; minutes <= 59; minutes++ {
			for seconds := 0; seconds <= 59; seconds++ {
				fmt.Print(hours, ":", minutes, ":", seconds, "\t")
			}
		}
	}
	fmt.Println()
}
