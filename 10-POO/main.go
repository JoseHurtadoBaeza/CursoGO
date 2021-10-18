// POO EN GO

package main

import (
	"fmt"
	"time" // Para manejar tipos fecha

	us "CursoGO/10-POO/user"
)

/*type usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}*/

// Herencia de la estructura usuario que podemos extender para dar polimorfismo
type pepe struct {
	// Esta es la manera de hacer herencia
	us.Usuario // Esto es como un puntero a Usuario
}

func main() {

	/*User := new(usuario)
	User.Id = 10
	User.Nombre = "Pablo"
	fmt.Println(User)*/

	u := new(pepe)
	u.AltaUsuario(1, "Pablo Tilotta", time.Now(), true)
	fmt.Println(u.Usuario)
}
