package main

import (
	"fmt"
	"strings"
)

func repeticiones2() {
	carros := []string{"Toyota", "Ford", "Chevrolet"}

	fmt.Println("Lista de carros inicial:", carros)
	fmt.Println("Ingresa nombres de carros. Si ingresas uno de la lista, el bucle se detendrá.")

	// Bucle infinito (simula un 'while true')
	for {
		var entrada string
		fmt.Print("\nIngresa una marca de carro: ")
		fmt.Scanln(&entrada)

		// Buscar si la entrada ya existe en la lista
		existe := false
		i := 0
		for i < len(carros) {
			if strings.EqualFold(carros[i], entrada) {
				existe = true
				break // Sale del bucle de búsqueda interno
			}
			i++
		}

		// Condición para detener el bucle principal
		if existe {
			fmt.Printf("¡'%s' ya está en la lista! Deteniendo el bucle...\n", entrada)
			break // Rompe el 'for' infinito (while true)
		}

		// Si no existe, lo agregamos y continuamos
		carros = append(carros, entrada)
		fmt.Println("Carro agregado. Lista actual:", carros)
	}

	fmt.Println("\nPrograma finalizado.")
}