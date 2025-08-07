package Entitys

type OrdenDeTrabajo struct {
	ID                int    `json:"id" db:"id"`
	Id_cliente        int    `json:"id_cliente" db:"id_cliente"`
	FechaCreacion     string `json:"fecha_creacion" db:"fecha_creacion"`
	FechaFinalizacion string `json:"fecha_finalizacion" db:"fecha_finalizacion"`
	Estado            string `json:"estado" db:"estado"`
	Descripcion       string `json:"descripcion" db:"descripcion"`
}
