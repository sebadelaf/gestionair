package Entitys

import "gorm.io/gorm"

// Protocolo define un conjunto de tareas a realizar.
type Protocolo struct {
	gorm.Model
	Nombre string `gorm:"size:255;not null;unique"`

	// Relación: Un Protocolo tiene muchas Tareas.
	Tareas []*Tarea `gorm:"foreignKey:ProtocoloID"`
}
