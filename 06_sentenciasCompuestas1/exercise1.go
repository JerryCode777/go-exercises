package main

import "fmt"

func main() {
	var m, n int
	fmt.Print("Ingrese el valor de m: ")
	fmt.Scan(&m)
	fmt.Print("Ingrese el valor de n: ")
	fmt.Scan(&n)

	for i := m; i <= n; i++ {
		if i%5 == 0 {
			fmt.Print(i, "\t")
		}
	}
	fmt.Println()
}
