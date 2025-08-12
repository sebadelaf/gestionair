package Entitys

import "gorm.io/gorm"

// Cliente representa a un cliente de la empresa.
type Cliente struct {
	gorm.Model
	Empresa   string `gorm:"size:255;not null"`
	Rut       string `gorm:"size:12;unique"`
	Contacto  string `gorm:"size:255"`
	Telefono  string `gorm:"size:20"`
	Email     string `gorm:"size:100"`
	Direccion string
	Notas     string

	// Relación: Un Cliente tiene muchos Equipos.
	Equipos []*Equipo `gorm:"foreignKey:ClienteID"`

	// Relación: Un Cliente tiene muchas Órdenes de Trabajo.
	OrdenesDeTrabajo []*OrdenDeTrabajo `gorm:"foreignKey:ClienteID"`
}
