package main

import (
	"fmt"
	"log"
)

func main() {
	arreglo := make([]string, 100)
	var i int

	for i < 100 {
		fmt.Println("Ingrese un numero de celular", i+1, " (q para finalizar): ")
		if _, err := fmt.Scan(&arreglo[i]); err != nil {
			log.Fatal(err)
		}
		if arreglo[i] == "q" {
			break
		}
		if len(arreglo[i]) != 9 {
			fmt.Println("No valido, debe ser de 9 digitos!")
			continue
		}
		fmt.Println("Ingresado!")
		i++
	}

	fmt.Println("Numeros de celular guardados: ")
	fmt.Println(arreglo)
}
