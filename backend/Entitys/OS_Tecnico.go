package Entitys

type OrdenDeServicioTecnico struct {
	ID        int `json:"id" db:"id"`
	IdOS      int `json:"id_os" db:"id_os"`
	IdTecnico int `json:"id_tecnico" db:"id_tecnico"`
}
