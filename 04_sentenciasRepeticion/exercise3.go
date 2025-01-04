package main

import "fmt"

func main() {
	var m, n int

	fmt.Print("Ingrese el valor de n: ")
	fmt.Scan(&m)
	fmt.Print("Ingrese el valor de m: ")
	fmt.Scan(&n)

	for i := m + 1; i < n; i = i + 2 {
		if m%2 == 0 {
			fmt.Println(i)
		} else {
			fmt.Println(i + 1)
		}
	}
}
