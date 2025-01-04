package main

import "fmt"

func main() {
	var numero1, numero2, module int
	fmt.Print("Ingrese el primer numero: ")
	fmt.Scan(&numero1)
	fmt.Print("Ingrese el segundo numero: ")
	fmt.Scan(&numero2)

	module = numero1 % numero2

	for module != 0 {
		numero1 = numero2
		numero2 = module
		module = numero1 % numero2
	}

	fmt.Println("El mcd es:", numero2)

}
