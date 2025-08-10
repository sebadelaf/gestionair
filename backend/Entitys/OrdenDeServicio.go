package Entitys

import (
	"gorm.io/gorm"
	"time"
)

// OrdenDeServicio es un trabajo específico sobre un equipo.
type OrdenDeServicio struct {
	gorm.Model
	OrdenDeTrabajoID uint   `gorm:"not null"`
	EquipoID         uint   `gorm:"not null"`
	ProtocoloID      uint   `gorm:"not null"`
	Estado           string `gorm:"size:50;not null;default:pendiente"`
	FechaServicio    time.Time
	Descripcion      string
	HorasTrabajadas  float32
	// Relación: Una Orden de Servicio es asignada a muchos Técnicos (Usuarios).
	TecnicosAsignados []*Usuario `gorm:"many2many:os_tecnicos_asignados;"`
}
