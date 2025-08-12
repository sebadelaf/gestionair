package Entitys

import (
	"gorm.io/gorm"
	"time"
)

// Equipo representa un equipo de aire acondicionado.
type Equipo struct {
	gorm.Model
	ClienteID        uint   `gorm:"not null"` // Clave foránea explícita
	Marca            string `gorm:"size:100"`
	Modelo           string `gorm:"size:100"`
	NumeroSerie      string `gorm:"size:100;unique"`
	TipoEquipo       string `gorm:"size:100"`
	Capacidad        string `gorm:"size:50"`
	Refrigerante     string `gorm:"size:50"`
	Tecnologia       string `gorm:"size:50"`
	Voltaje          string `gorm:"size:50"`
	FechaInstalacion time.Time
	Ubicacion        string

	// Relación: Un Equipo puede estar en muchas Órdenes de Trabajo.
	OrdenesDeTrabajo []*OrdenDeTrabajo `gorm:"many2many:ot_equipos;"`
}
