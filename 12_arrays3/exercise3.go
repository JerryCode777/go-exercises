package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	horas := make([]int, 31)
	for i := 0; i < len(horas); i++ {
		horas[i] = rand.Intn(9)
	}
	fmt.Println(horas)

	var dias, ultimo int
	fmt.Print("Ingrese la cantidad de dias que pasen: ")
	if _, err := fmt.Scan(&dias); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < dias; i++ {
		fmt.Print("Ingrese la cantidad de horas a trabajar en el dia 30: ")
		if _, err := fmt.Scan(&ultimo); err != nil {
			log.Fatal(err)
		}
		for j := 0; j < len(horas)-1; j++ {
			horas[j] = horas[j+1]
		}
		horas[len(horas)-1] = ultimo
		fmt.Println(horas)
	}
}
