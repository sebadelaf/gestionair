package Entitys

type OrdenTrabajoTecnico struct {
	ID        int `json:"id" db:"id"`
	IdOT      int `json:"id_ot" db:"id_ot"`
	IdTecnico int `json:"id_tecnico" db:"id_tecnico"`
}
