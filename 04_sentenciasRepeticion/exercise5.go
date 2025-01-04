package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Ingresa el valor de n: ")
	fmt.Scan(&n)

	fmt.Println("nro", "\t", "^2")

	for i := 1; i <= n; i++ {
		fmt.Println(i, "\t", math.Pow(float64(i), 2))
	}
}
