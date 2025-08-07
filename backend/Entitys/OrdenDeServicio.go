package Entitys

type OrdenDeServicio struct {
	ID                  int    `json:"id" db:"id"`
	Id_orden_de_trabajo int    `json:"id_orden_de_trabajo" db:"id_orden_de_trabajo"`
	FechaCreacion       string `json:"fecha_creacion" db:"fecha_creacion"`
	FechaFinalizacion   string `json:"fecha_finalizacion" db:"fecha_finalizacion"`
	Estado              string `json:"estado" db:"estado"`
	Descripcion         string `json:"descripcion" db:"descripcion"`
}
