// MANEJO DE ARCHIVOS EN GO

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	leoArchivo()
	leoArchivo2()
	graboArchivo()
}

func leoArchivo() {
	// Este es más fácil, pero tiene una limitación,
	// me lee el archivo de un sólo intento, pero no me lo deja manipular mucho
	archivo, err := ioutil.ReadFile("./archivo.txt")

	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(archivo))
	}

}

func leoArchivo2() {
	// En cambio, usando el paquete del sistema operativo voy a abrir el archivo
	// y con la función scanner vamos a poder recorrer de forma secuencial
	// registro por registro nuestro archivo
	archivo, err := os.Open("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		scanner := bufio.NewScanner(archivo)
		// Lee hasta que una instrucción sea false
		for scanner.Scan() {
			registro := scanner.Text() // Convierte la línea leída que se encuentra en el buffer, en una cadena de texto
			fmt.Printf("Linea > " + registro + "\n")
		}
	}

	archivo.Close()

}

func graboArchivo() {
	archivo, err := os.Create
}
