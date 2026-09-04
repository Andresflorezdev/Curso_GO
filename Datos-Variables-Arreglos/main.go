package main

import "fmt"
func main()  {
	var nombre string = "Andrews";
	var apellido = "Paternina"
	segundoNombre := "Felipe"
	fmt.Println(nombre + segundoNombre + apellido)
	/*fmt.Println(apellido)
	fmt.Println(segundoNombre)/*

	/* Parte numerica */
	var numero int16 = 2026
	var reducido int8 = 127
	edad := 19
	fmt.Println(numero + int16(reducido) + int16(edad))
	/*fmt.Println(reducido)
	fmt.Println(edad)*/

	/* Arreglos */
	var listaFrutas = [4]string {"pera","manzana", "piña"}
	fmt.Println(listaFrutas[1])

	/*listaPaises := [3]string{"argelia", "peru", "honduras"}*/
	listaPaises := []string{"argelia", "peru", "honduras"}
	fmt.Println(listaPaises)
	/*listaPaises[0] = "colombia"*/
	listaPaises = append(listaPaises, "colombia")
	fmt.Println(listaPaises)

	listaPaises2 := listaPaises[1:3]
	fmt.Println(listaPaises2)

	listaPaises3 := listaPaises[2:]
	fmt.Println(listaPaises3)
}
