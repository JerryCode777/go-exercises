package main

import "fmt"

func main() {
	fmt.Println("  ", 5, "  ", 10, " ", 15, " ", 20, " ", 25)

	for j := 11; j <= 20; j = j + 1 {
		fmt.Print(j, " ")
		for i := 5; i <= 25; i = i + 5 {
			fmt.Print(i*j*19, " ")
		}
		fmt.Println()
	}
}
