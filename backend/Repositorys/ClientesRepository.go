package Repositorys

import (
	"github.com/sebadelaf/gestionair/Entitys"

	"gorm.io/gorm"
)

type ClienteRepository interface {
	Create(cliente *Entitys.Cliente) error
	FindByID(id uint) (*Entitys.Cliente, error)
	FindAll() ([]*Entitys.Cliente, error)
	Update(cliente *Entitys.Cliente) error
	Delete(id uint) error
}

type clienteRepository struct {
	db *gorm.DB // Mantiene una referencia a la conexión de la base de datos (GORM).
}

//CRUD

func NewClienteRepository(db *gorm.DB) ClienteRepository {
	return &clienteRepository{db: db}
}

// NewClienteRepository crea una nueva instancia de ClienteRepository con la conexión a la base de datos proporcionada.
func (r *clienteRepository) Create(cliente *Entitys.Cliente) error {
	return r.db.Create(cliente).Error
}

// FindById busca un cliente por su ID y devuelve un puntero al cliente encontrado o un error si no se encuentra.
func (r *clienteRepository) FindByID(id uint) (*Entitys.Cliente, error) {
	var cliente Entitys.Cliente
	err := r.db.Preload("Equipos").First(&cliente, id).Error
	return &cliente, err
}

// FindAll devuelve una lista de todos los clientes en la base de datos.
func (r *clienteRepository) FindAll() ([]*Entitys.Cliente, error) {
	var clientes []*Entitys.Cliente
	err := r.db.Find(&clientes).Error
	return clientes, err
}

// Update actualiza un cliente existente en la base de datos.
func (r *clienteRepository) Update(cliente *Entitys.Cliente) error {
	return r.db.Save(cliente).Error
}

// Delete elimina un cliente de la base de datos por su ID.
func (r *clienteRepository) Delete(id uint) error {
	return r.db.Delete(&Entitys.Cliente{}, id).Error
}
