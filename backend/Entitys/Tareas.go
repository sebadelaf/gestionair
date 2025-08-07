package Entitys

type Tareas struct {
	IdTarea      int    `json:"id_tarea" db:"id_tarea"`
	Id_protocolo int    `json:"id_protocolo" db:"id_protocolo"`
	Tarea        string `json:"tarea" db:"tarea"`
}
