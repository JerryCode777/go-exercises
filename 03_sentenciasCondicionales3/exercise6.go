package main

import "fmt"

func main() {
	var edad, antiguedad int
	fmt.Print("Ingrese la edad de la persona: ")
	fmt.Scan(&edad)
	fmt.Print("Ingrese la antiguedad: ")
	fmt.Scan(&antiguedad)

	switch {
	case edad >= 60 && antiguedad < 25:
		fmt.Println("jubilacion por edad")
	case edad < 60 && antiguedad >= 25:
		fmt.Println("jubilacion por antiguedad joven")
	case edad >= 60 && antiguedad >= 25:
		fmt.Println("jubilacion por antiguedad adulta")
	default:
		fmt.Println("No accede a jubilacion")
	}
}
