package Repositorys

import (
	"github.com/sebadelaf/gestionair/Entitys"

	"gorm.io/gorm"
)

type ProtocoloRepository interface {
	Create(protocolo *Entitys.Protocolo) error
	FindByID(id uint) (*Entitys.Protocolo, error)
	FindAll() ([]*Entitys.Protocolo, error)
	Update(protocolo *Entitys.Protocolo) error
	Delete(id uint) error
	FindByNombre(nombre string) (*Entitys.Protocolo, error)
}

type protocoloRepository struct {
	db *gorm.DB // Mantiene una referencia a la conexión de la base de datos (GORM).
}

// NewProtocoloRepository crea una nueva instancia de ProtocoloRepository con la conexión a la base de datos proporcionada.
func NewProtocoloRepository(db *gorm.DB) ProtocoloRepository {
	return &protocoloRepository{db: db}
}

// Create inserta un nuevo protocolo en la base de datos.
func (r *protocoloRepository) Create(protocolo *Entitys.Protocolo) error {
	return r.db.Create(protocolo).Error
}

// FindByID busca un protocolo por su ID y devuelve un puntero al protocolo encontrado o un error si no se encuentra.
func (r *protocoloRepository) FindByID(id uint) (*Entitys.Protocolo, error) {
	var protocolo Entitys.Protocolo
	err := r.db.Preload("Tareas").First(&protocolo, id).Error
	return &protocolo, err
}

// FindAll devuelve una lista de todos los protocolos en la base de datos.
func (r *protocoloRepository) FindAll() ([]*Entitys.Protocolo, error) {
	var protocolos []*Entitys.Protocolo
	err := r.db.Preload("Tareas").Find(&protocolos).Error
	return protocolos, err
}

// Update actualiza un protocolo existente en la base de datos.
func (r *protocoloRepository) Update(protocolo *Entitys.Protocolo) error {
	return r.db.Save(protocolo).Error
}

// Delete elimina un protocolo de la base de datos por su ID.
func (r *protocoloRepository) Delete(id uint) error {
	return r.db.Delete(&Entitys.Protocolo{}, id).Error
}

// FindByNombre busca un protocolo por su nombre y devuelve un puntero al protocolo encontrado o un error si no se encuentra.
func (r *protocoloRepository) FindByNombre(nombre string) (*Entitys.Protocolo, error) {
	var protocolo Entitys.Protocolo
	err := r.db.Where("nombre = ?", nombre).First(&protocolo).Error
	if err != nil {
		return nil, err
	}
	return &protocolo, nil
}
