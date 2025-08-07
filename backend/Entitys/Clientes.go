package Entitys

type Clientes struct {
	ID            int    `json:"id" db:"id"`
	Nombre        string `json:"nombre" db:"nombre"`
	Direccion     string `json:"direccion" db:"direccion"`
	Telefono      string `json:"telefono" db:"telefono"`
	Email         string `json:"email" db:"email"`
	Lista_Equipos []int  `json:"lista_equipos"` // Lista de IDs de equipos asociados al cliente
	Lista_ordenes []int  `json:"lista_ordenes"` // Lista de IDs de órdenes asociadas al cliente
}
