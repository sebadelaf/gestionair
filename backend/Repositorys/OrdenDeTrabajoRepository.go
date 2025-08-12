package Repositorys

import (
	"github.com/sebadelaf/gestionair/Entitys"

	"gorm.io/gorm"
)

type OrdenDeTrabajoRepository interface {
	Create(orden *Entitys.OrdenDeTrabajo) error
	FindByID(id uint) (*Entitys.OrdenDeTrabajo, error)
	FindAll() ([]*Entitys.OrdenDeTrabajo, error)
	Update(orden *Entitys.OrdenDeTrabajo) error
	Delete(id uint) error
	FindByEquipoID(equipoID uint) ([]*Entitys.OrdenDeTrabajo, error)
	FindByEstado(estado string) ([]*Entitys.OrdenDeTrabajo, error)
	FindByClienteID(clienteID uint) ([]*Entitys.OrdenDeTrabajo, error)
}
type ordenDeTrabajoRepository struct {
	db *gorm.DB // Mantiene una referencia a la conexión de la base de datos (GORM).
}

// NewOrdenDeTrabajoRepository crea una nueva instancia de OrdenDeTrabajoRepository con la conexión a la base de datos proporcionada.
func NewOrdenDeTrabajoRepository(db *gorm.DB) OrdenDeTrabajoRepository {
	return &ordenDeTrabajoRepository{db: db}
}

// Create inserta una nueva orden de trabajo en la base de datos.
func (r *ordenDeTrabajoRepository) Create(orden *Entitys.OrdenDeTrabajo) error {
	return r.db.Create(orden).Error
}

// FindByID busca una orden de trabajo por su ID y devuelve un puntero a la orden encontrada o un error si no se encuentra.
func (r *ordenDeTrabajoRepository) FindByID(id uint) (*Entitys.OrdenDeTrabajo, error) {
	var orden Entitys.OrdenDeTrabajo
	err := r.db.Preload("Equipo").Preload("Cliente").First(&orden, id).Error
	return &orden, err
}

// FindAll devuelve una lista de todas las órdenes de trabajo en la base de datos.
func (r *ordenDeTrabajoRepository) FindAll() ([]*Entitys.OrdenDeTrabajo, error) {
	var ordenes []*Entitys.OrdenDeTrabajo
	err := r.db.Preload("Equipo").Preload("Cliente").Find(&ordenes).Error
	return ordenes, err
}

// Update actualiza una orden de trabajo existente en la base de datos.
func (r *ordenDeTrabajoRepository) Update(orden *Entitys.OrdenDeTrabajo) error {
	return r.db.Save(orden).Error
}

// Delete elimina una orden de trabajo de la base de datos por su ID.
func (r *ordenDeTrabajoRepository) Delete(id uint) error {
	return r.db.Delete(&Entitys.OrdenDeTrabajo{}, id).Error
}

// FindByEquipoID busca todas las órdenes de trabajo asociadas a un equipo específico por su ID y devuelve una lista de órdenes.
func (r *ordenDeTrabajoRepository) FindByEquipoID(equipoID uint) ([]*Entitys.OrdenDeTrabajo, error) {
	var ordenes []*Entitys.OrdenDeTrabajo
	err := r.db.Where("equipo_id = ?", equipoID).Preload("Equipo").Preload("Cliente").Find(&ordenes).Error
	return ordenes, err
}

// FindByEstado busca todas las órdenes de trabajo por su estado y devuelve una lista de órdenes.
func (r *ordenDeTrabajoRepository) FindByEstado(estado string) ([]*Entitys.OrdenDeTrabajo, error) {
	var ordenes []*Entitys.OrdenDeTrabajo
	err := r.db.Where("estado = ?", estado).Preload("Equipo").Preload("Cliente").Find(&ordenes).Error
	return ordenes, err
}

// FindByClienteID busca todas las órdenes de trabajo asociadas a un cliente específico por su ID y devuelve una lista de órdenes.
func (r *ordenDeTrabajoRepository) FindByClienteID(clienteID uint) ([]*Entitys.OrdenDeTrabajo, error) {
	var ordenes []*Entitys.OrdenDeTrabajo
	err := r.db.Where("cliente_id = ?", clienteID).Preload("Equipo").Preload("Cliente").Find(&ordenes).Error
	return ordenes, err
}
