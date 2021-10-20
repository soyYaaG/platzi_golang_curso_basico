package main

import "fmt"

func main() {
	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("Constantes")
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)
	fmt.Println("")

	// Declaración de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println("Variables")
	fmt.Println(base, altura, area)
	fmt.Println("")

	// Zero values
	var nombre string
	var edad int
	var peso float64
	var estudiante bool

	fmt.Println("Nombre:", nombre)
	fmt.Println("Edad:", edad)
	fmt.Println("Peso:", peso)
	fmt.Println("Estudiante:", estudiante)
	fmt.Println("")

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)
}