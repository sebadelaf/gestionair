package Entitys

type Usuarios struct {
	ID       int    `json:"id" db:"id"`
	Nombre   string `json:"nombre" db:"nombre"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
	Telefono string `json:"telefono" db:"telefono"`
	Rol      string `json:"rol" db:"rol"` // Puede ser "admin", "tecnico", etc.
}
