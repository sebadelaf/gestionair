package Entitys

import (
	"time"

	"gorm.io/gorm"
)

// OrdenDeTrabajo es la solicitud general de un cliente.
type OrdenDeTrabajo struct {
	gorm.Model
	ClienteID        uint   `gorm:"not null"`
	UsuarioCreadorID uint   `gorm:"not null"`
	Estado           string `gorm:"size:50;not null;default:pendiente"`
	FechaCierre      time.Time
	Descripcion      string

	// Relación: Una Orden de Trabajo tiene muchas Órdenes de Servicio.
	OrdenesDeServicio []*OrdenDeServicio `gorm:"foreignKey:OrdenDeTrabajoID"`
	// Relación: Una Orden de Trabajo puede involucrar muchos Equipos.
	Equipos []*Equipo `gorm:"many2many:ot_equipos;"`
}
