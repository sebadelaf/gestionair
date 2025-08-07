package Entitys

type OrdenTrabajoEquipo struct {
	ID       string `json:"id" db:"id"`
	IdOT     int    `json:"id_ot" db:"id_ot"`
	IdEquipo int    `json:"id_equipo" db:"id_equipo"`
}
