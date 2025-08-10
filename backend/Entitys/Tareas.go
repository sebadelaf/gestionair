package Entitys

import "gorm.io/gorm"

// Tarea representa una única tarea dentro de un checklist de un protocolo.
type Tarea struct {
	gorm.Model
	ProtocoloID uint   `gorm:"not null"` // Clave foránea para el protocolo
	Descripcion string `gorm:"not null"`
	Obligatorio bool
}
