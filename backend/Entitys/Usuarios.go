package Entitys

import "gorm.io/gorm"

// Usuario representa a una persona que puede interactuar con el sistema.
type Usuario struct {
	gorm.Model
	Nombre       string `gorm:"size:255;not null"`
	Email        string `gorm:"size:100;not null;unique"`
	PasswordHash string `gorm:"not null"`
	Rol          string `gorm:"size:50;not null"`
	Telefono     string `gorm:"size:20"`
	// Relación: Un Usuario (técnico) puede ser asignado a muchas Órdenes de Servicio.
	OrdenesDeServicioAsignadas []OrdenDeServicio `gorm:"many2many:os_tecnicos_asignados;"`
}
