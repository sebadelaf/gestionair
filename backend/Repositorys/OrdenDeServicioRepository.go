package Repositorys

import (
	"github.com/sebadelaf/gestionair/Entitys"

	"gorm.io/gorm"
)

type OrdenDeServicioRepository interface {
	Create(orden *Entitys.OrdenDeServicio) error
	FindByID(id uint) (*Entitys.OrdenDeServicio, error)
	FindAll() ([]*Entitys.OrdenDeServicio, error)
	Update(orden *Entitys.OrdenDeServicio) error
	Delete(id uint) error
	FindByOrdenID(ordenID uint) (*Entitys.OrdenDeServicio, error)
}
type ordenDeServicioRepository struct {
	db *gorm.DB // Mantiene una referencia a la conexión de la base de datos (GORM).
}

// NewOrdenDeServicioRepository crea una nueva instancia de OrdenDeServicioRepository con la conexión a la base de datos proporcionada.
func NewOrdenDeServicioRepository(db *gorm.DB) OrdenDeServicioRepository {
	return &ordenDeServicioRepository{db: db}
}

// Create inserta una nueva orden de servicio en la base de datos.
func (r *ordenDeServicioRepository) Create(orden *Entitys.OrdenDeServicio) error {
	return r.db.Create(orden).Error
}

// FindByID busca una orden de servicio por su ID y devuelve un puntero a la orden encontrada o un error si no se encuentra.
func (r *ordenDeServicioRepository) FindByID(id uint) (*Entitys.OrdenDeServicio, error) {
	var orden Entitys.OrdenDeServicio
	err := r.db.Preload("Cliente").First(&orden, id).Error
	return &orden, err
}

// FindAll devuelve una lista de todas las órdenes de servicio en la base de datos.
func (r *ordenDeServicioRepository) FindAll() ([]*Entitys.OrdenDeServicio, error) {
	var ordenes []*Entitys.OrdenDeServicio
	err := r.db.Preload("Cliente").Find(&ordenes).Error
	return ordenes, err
}

// Update actualiza una orden de servicio existente en la base de datos.
func (r *ordenDeServicioRepository) Update(orden *Entitys.OrdenDeServicio) error {
	return r.db.Save(orden).Error
}

// Delete elimina una orden de servicio de la base de datos por su ID.
func (r *ordenDeServicioRepository) Delete(id uint) error {
	return r.db.Delete(&Entitys.OrdenDeServicio{}, id).Error
}

// FindByOrdenID busca una orden de servicio por su ID de orden y devuelve un puntero a la orden encontrada o un error si no se encuentra.
func (r *ordenDeServicioRepository) FindByOrdenID(ordenID uint) (*Entitys.OrdenDeServicio, error) {
	var orden Entitys.OrdenDeServicio
	err := r.db.Where("orden_de_trabajo_id = ?", ordenID).Preload("Cliente").First(&orden).Error
	if err != nil {
		return nil, err
	}
	return &orden, nil
}
