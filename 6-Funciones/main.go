// FUNCIONES EN GO

package main

import "fmt"

func main() {

	/*fmt.Println(uno(5))

	numero, estado := dos(2)

	fmt.Println(numero)
	fmt.Println(estado)*/

	/*fmt.Println(Calculo(5, 46))
	fmt.Println(Calculo(5, 23, 45, 68))
	fmt.Println(Calculo(5))
	fmt.Println(Calculo(5, 46, 17, 25, 26, 98, 17))*/

	fmt.Println(CalculoVariante(5, 46))
	fmt.Println(CalculoVariante(5, 23, 45, 68))
	fmt.Println(CalculoVariante(5))
	fmt.Println(CalculoVariante(5, 46, 17, 25, 26, 98, 17))

}

/*func uno(numero int) int {
	return numero * 2
}*/

// Variante de la función uno en la que indicamos el nombre de la variable de salida
func uno(numero int) (z int) {
	z = numero * 2
	return
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

// Ejemplo con parámetros variables o dinámicos, es decir, no sabemos el número de parámetros de entrada
// Aquí le estamos diciendo, vas a recibir parámetros enteros pero no te puedo indicar cuantos
func Calculo(numero ...int) int {
	total := 0
	// La función range me devuelve siempre 2 valores: el número de elemento y , lo que hace es tomar un rango de parámetros
	// lo que se coloca a la derecha tiene que ser una lista, un vector, etc.
	// En este caso, al ser parámetros variables los convierte en una lista en memoria y
	// entonces con esta función podemos iterar por cada uno de sus elementos.
	// En GO cuando yo tengo una función o una instrucción que me devuelve algún tipo de dato que no quiero usar,
	// se puede usar el _ para alojar esa variable que no la voy a usar. Evitando que ocupe en memoria.
	// En este caso, no nos interesa saber el número del elemento.
	// Y num si es una variable nueva donde voy a ir guardando los números
	for _, num := range numero {
		total = total + num
	}

	return total
}

// Variante de la anterior donde no usamos el _
func CalculoVariante(numero ...int) int {
	total := 0
	for item, num := range numero {
		total = total + num
		fmt.Printf("\n Item %d \n", item)
	}

	return total
}
