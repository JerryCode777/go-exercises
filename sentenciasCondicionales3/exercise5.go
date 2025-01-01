package main

import "fmt"

func main() {
	var color string
	var descuento, monto float64

	fmt.Print("Ingrese el color de la bolita: ")
	fmt.Scan(&color)

	fmt.Print("Ingrese el monto: ")
	fmt.Scan(&monto)

	switch color {
	case "Blanco":
		descuento = 0
	case "Verde":
		descuento = 10
	case "Amarillo":
		descuento = 25
	case "Azul":
		descuento = 50
	case "Rojo":
		descuento = 100
	default:
		fmt.Println("Ingrese una opcion valida de descuento!")
		return
	}

	montoFinal := monto - descuento*monto/100

	fmt.Println("El monto final que debe pagar es: ", montoFinal)

}
