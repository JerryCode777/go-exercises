package main

import "fmt"

func muestra(a, b float64) {
	if b < 0 {
		fmt.Printf("%.2f%.2fi\n", a, b)
		return
	}
	fmt.Printf("%.2f+%.2fi\n", a, b)
}

func main() {
	var a, b float64
	fmt.Print("Ingrese el valor de la parte real: ")
	fmt.Scan(&a)
	fmt.Print("Ingrese el valor de la parte imaginaria: ")
	fmt.Scan(&b)

	muestra(a, b)

}
