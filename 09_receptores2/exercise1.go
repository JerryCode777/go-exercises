package main

import "fmt"

func muestra(a, b float64) {
	fmt.Printf("(%.2f,%.2f)\n", a, b)
}
func main() {
	//programa que muestre un par ordenado con el formato a,b metodo muestra
	var a, b float64
	fmt.Print("Ingrese el valor de a: ")
	fmt.Scan(&a)
	fmt.Print("Ingrese el valor de b: ")
	fmt.Scan(&b)

	muestra(a, b)
}
