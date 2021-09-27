package user

import "time"

type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

// Se indica en la declaración de la función que cuando dentro de esta se haga referencia a this
// vamos a estar apuntando a la estructura Usuario
func (this *Usuario) AltaUsuario(id int, nombre string, fechaalta time.Time, status bool) {

	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaalta
	this.Status = status

}
