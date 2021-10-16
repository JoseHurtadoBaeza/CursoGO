// MAPAS EN GO
package main

import "fmt"

func main() {
	// En este caso le estamos diciendo que la clave va a ser de tipo string y el valor asociado también
	/*paises := make(map[string]string, 5)
	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises["Mexico"])
	fmt.Println(paises)*/

	// Los elementos los guarda en memoria alfabéticamente
	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	campeonato["River Plate"] = 25    // Añadimos un item al mapa
	campeonato["Chivas"] = 55         // Modificamos un valor
	delete(campeonato, "Real Madrid") // Borrar un elemento

	fmt.Println(campeonato)

	// Vamos a recorrer un mapa con range
	for equipo, puntaje := range campeonato { // El range devuelve dos valores: clave y valor
		fmt.Printf("El equipo %s, tiene un puntaje de : %d \n", equipo, puntaje)
	}

	// Consultar si un elemento existe en un mapa
	puntaje, existe := campeonato["Mineiro"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)

	puntaje, existe = campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)

}
