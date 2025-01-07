package main

import "fmt"

// definimos un tipo struct
type Persona struct {
	Nombre string
}

// metodo asociado al tipo de persona utilizando un receptor de valor
func (p Persona) saludar() {
	fmt.Println("Hola mi nombre es: ", p.Nombre)
}

func main() {
	Personita := Persona{Nombre: "Juan"}
	Personita.saludar() //llamamos al metodo saludar, muy parecido a POO
}
