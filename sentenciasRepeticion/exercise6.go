package main

import "fmt"

func main() {
	var n int
	fmt.Print("Ingrese el valor de n: ")
	fmt.Scan(&n)

	fmt.Println("Numeros pares <= que n")
	for i := 0; i <= n; i = i + 2 {
		fmt.Println(i)
	}
}
