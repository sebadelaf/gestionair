package Entitys

type OrdenDeTrabajo struct {
	ID                int    `json:"id" db:"id"`
	FechaCreacion     string `json:"fecha_creacion" db:"fecha_creacion"`
	FechaFinalizacion string `json:"fecha_finalizacion" db:"fecha_finalizacion"`
	Estado            string `json:"estado" db:"estado"`
	Descripcion       string `json:"descripcion" db:"descripcion"`
	Lista_Equipos     []int  `json:"lista_equipos"`
	Lista_ordenes     []int  `json:"lista_ordenes"`
	Lista_tecnicos    []int  `json:"lista_tecnicos"`
}
