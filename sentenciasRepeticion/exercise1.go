package main

import "fmt"

func main() {
	//programa que muestre los n primeros numeros no negativos
	var n int
	fmt.Print("Ingrese el valor de n: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Println(i, "\t")
	}
}
