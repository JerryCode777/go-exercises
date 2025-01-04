package main

import "fmt"

func main() {
	var entero int
	fmt.Print("Ingresar el numero entero: ")
	fmt.Scan(&entero)

	for i := 1; i <= 9; i++ {
		fmt.Println(entero, " x ", i, " = ", entero*i)
	}
}
