package main

import "fmt"

func main() {
	var numero1, numero2 int

	fmt.Print("Ingrese el primer numero entero: ")
	fmt.Scan(&numero1)

	fmt.Print("Ingrese el segundo numero entero: ")
	fmt.Scan(&numero2)

	if numero1 > numero2 {
		fmt.Println(numero2, "-", numero1)
	} else {
		fmt.Println(numero1, "-", numero2)
	}
}
