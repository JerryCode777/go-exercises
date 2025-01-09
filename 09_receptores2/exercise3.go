package main

import (
	"fmt"
	"math"
)

func muestraHipotenusa(a, b float64) {
	hipotenusa := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	fmt.Println("La hipotenusa es: ", hipotenusa)
}

func hipotenusa(a, b float64) float64 {
	hipotenusa := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	return hipotenusa
}

func main() {
	var cateto1, cateto2 float64
	fmt.Println("Ingrese el primer cateto: ")
	fmt.Scan(&cateto1)
	fmt.Println("Ingrese el segundo cateto: ")
	fmt.Scan(&cateto2)

	muestraHipotenusa(cateto1, cateto2)
	fmt.Println("retorna: ", hipotenusa(cateto1, cateto2))
}
