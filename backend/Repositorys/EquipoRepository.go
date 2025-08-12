package Repositorys

import (
	"github.com/sebadelaf/gestionair/Entitys"

	"gorm.io/gorm"
)

type EquipoRepository interface {
	Create(equipo *Entitys.Equipo) error
	FindByID(id uint) (*Entitys.Equipo, error)
	FindAll() ([]*Entitys.Equipo, error)
	Update(equipo *Entitys.Equipo) error
	Delete(id uint) error
	FindByClienteID(clienteID uint) ([]*Entitys.Equipo, error)
}

type equipoRepository struct {
	db *gorm.DB // Mantiene una referencia a la conexión de la base de datos (GORM).
}

// NewEquipoRepository crea una nueva instancia de EquipoRepository con la conexión a la base de datos proporcionada.
func NewEquipoRepository(db *gorm.DB) EquipoRepository {
	return &equipoRepository{db: db}
}

// CRUD
func (r *equipoRepository) Create(equipo *Entitys.Equipo) error {
	return r.db.Create(equipo).Error
}
func (r *equipoRepository) FindByID(id uint) (*Entitys.Equipo, error) {
	var equipo Entitys.Equipo
	err := r.db.Preload("OrdenesDeTrabajo").First(&equipo, id).Error
	return &equipo, err
}
func (r *equipoRepository) FindAll() ([]*Entitys.Equipo, error) {
	var equipos []*Entitys.Equipo
	err := r.db.Find(&equipos).Error
	return equipos, err
}
func (r *equipoRepository) Update(equipo *Entitys.Equipo) error {
	return r.db.Save(equipo).Error
}

func (r *equipoRepository) Delete(id uint) error {
	return r.db.Delete(&Entitys.Equipo{}, id).Error
}
func (r *equipoRepository) FindByClienteID(clienteID uint) ([]*Entitys.Equipo, error) {
	var equipos []*Entitys.Equipo
	err := r.db.Where("cliente_id = ?", clienteID).Find(&equipos).Error
	return equipos, err
}
