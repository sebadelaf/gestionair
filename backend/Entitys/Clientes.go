package Entitys

type Clientes struct {
	ID        int    `json:"id" db:"id"`
	Nombre    string `json:"nombre" db:"nombre"`
	Direccion string `json:"direccion" db:"direccion"`
	Telefono  string `json:"telefono" db:"telefono"`
	Email     string `json:"email" db:"email"`
}
