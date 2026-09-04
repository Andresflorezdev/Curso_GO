package main

import "fmt"

func mapas()  {
	miMapa := map[string]string {
		"medellin":"antioquia",
		"bogota":"cundinamarca",
		"manizales":"caldas",
	}
	/*fmt.Println("mapa de paises", miMapa)*/
	fmt.Println("medellin",miMapa["medellin"])

	miMapa["Rionegro"] = "Guarne"
	fmt.Println("mapa de paises", miMapa)

	delete(miMapa, "Rionegro")
	fmt.Println("mapa de paises", miMapa)
}
