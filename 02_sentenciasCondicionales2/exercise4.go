package main

import "fmt"

func main() {
	var tasaInteres, capital float64
	fmt.Print("Ingrese la tasa de interes anual (%): ")
	fmt.Scan(&tasaInteres)
	fmt.Print("Ingrese el capital: ")
	fmt.Scan(&capital)

	interes := capital * tasaInteres / 100

	if interes > 200 {
		fmt.Println("Se recomienda invertir para poder tener ", interes+capital, " al final del anio")
	} else {
		fmt.Println("No es rentable invertir")
	}
}
