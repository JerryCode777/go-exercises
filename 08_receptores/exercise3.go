package main

import (
	"fmt"
	"math"
)

func distancia(Tx, Ty, Sx, Sy float64) float64 {
	result := math.Sqrt(math.Pow(Tx-Sx, 2) + math.Pow(Ty-Sy, 2))
	return result
}

func main() {
	//programa que calcule la distancia entre strabucks y la ubicacion del turista
	var Tx, Ty, Sx, Sy float64
	fmt.Print("Ingrese la coordenada x del turista:")
	if _, err := fmt.Scan(&Tx); err != nil {
		fmt.Errorf("error")
		return
	}
	fmt.Print("Ingrese la coordenada y del turista:")
	if _, err := fmt.Scan(&Ty); err != nil {
		fmt.Errorf("error")
		return
	}
	fmt.Print("Ingrese la coordenada x del Starbucks:")
	if _, err := fmt.Scan(&Sx); err != nil {
		fmt.Errorf("error")
		return
	}
	fmt.Print("Ingrese la coordenada y del Starbucks:")
	if _, err := fmt.Scan(&Sy); err != nil {
		fmt.Errorf("error")
		return
	}

	fmt.Println("la distancia es: ", distancia(Tx, Ty, Sx, Sy))

}
