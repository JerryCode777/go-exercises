package main

import "fmt"

func main() {
	//mostrar todas las combinaciones posibles de 3 bits
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				fmt.Print(i)
				fmt.Print(j)
				fmt.Print(k, "\t")
			}
		}
	}
	fmt.Println()
}
