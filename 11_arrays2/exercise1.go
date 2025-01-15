package main

import (
	"fmt"
	"log"
)

func main() {
	var n int
	fmt.Print("Ingrese el numero de notas: ")
	if _, err := fmt.Scan(&n); err != nil {
		log.Fatal(err)
	}

	arreglo := make([]int, n)

	fmt.Println("Ingrese las notas: ")
	for i := 0; i < n; i++ {
		fmt.Print("Ingrese la nota ", i+1, ": ")
		if _, err := fmt.Scan(&arreglo[i]); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Notas:")
	fmt.Println(arreglo)

}
