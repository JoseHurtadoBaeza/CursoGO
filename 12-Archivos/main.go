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
	graboArchivo2()
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

// Crea un archivo nuevo, me va a borrar el contenido y vamos a poder guardar contenido nuevo
func graboArchivo() {
	// Create es lo mismo que un Open, sólamente que me lo abre de output (el archivo), me lo borra si ya teníamos un archivo con ese nombre (cuidado)
	// y lo que hace luego es regrabar todo lo que tiene en las líneas de código inferior en dicho archivo.
	archivo, err := os.Create("./nuevoArchivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}

	// Println con la diferencia de que ahora se graba en un archivo
	fmt.Fprintln(archivo, "Esta es una línea nueva")
	archivo.Close() // El create al igual que el Open necesita de un Close

}

// Aquí en cambio, vamos a ver como añadir texto a un archivo que ya existe
func graboArchivo2() {
	fileName := "./nuevoArchivo.txt"
	// El Append me va a agregar al final del archivo una nueva línea
	if Append(fileName, "\nEsta es una segunda línea") == false {
		fmt.Println("Error en la 2da linea")
	}

}

func Append(archivo string, texto string) bool {
	// La constante WRONLY es para indicar que el archivo es para leer y escrbir
	// La constante APPEND indica que queremos usar el modo que no limpia el archivo, no me lo blanquea,
	// sino que lo va a abrir para poder escribir al final y agragarle registros
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644) // Este número del final es para indicar los permisos del archivo

	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}

	// La función WriteString es para grabar un String dentro de "archivo"
	_, err = arch.WriteString(texto) // No nos interesa lo que devuelve la función, por eso usamos "_"

	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}

	return true

}
