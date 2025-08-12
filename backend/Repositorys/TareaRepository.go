package Repositorys

import (
	"github.com/sebadelaf/gestionair/Entitys"

	"gorm.io/gorm"
)

// TareaRepository define las operaciones CRUD para la entidad Tarea.
type TareaRepository interface {
	Create(tarea *Entitys.Tarea) error
	FindByID(id uint) (*Entitys.Tarea, error)
	FindAll() ([]*Entitys.Tarea, error)
	Update(tarea *Entitys.Tarea) error
	Delete(id uint) error
	FindByProtocoloID(protocoloID uint) ([]*Entitys.Tarea, error)
}
type tareaRepository struct {
	db *gorm.DB // Mantiene una referencia a la conexión de la base de datos (GORM).
}

// NewTareaRepository crea una nueva instancia de TareaRepository con la conexión a la base de datos proporcionada.
func NewTareaRepository(db *gorm.DB) TareaRepository {
	return &tareaRepository{db: db}
}

// Create inserta una nueva tarea en la base de datos.
func (r *tareaRepository) Create(tarea *Entitys.Tarea) error {
	return r.db.Create(tarea).Error
}

// FindByID busca una tarea por su ID y devuelve un puntero a la tarea encontrada o un error si no se encuentra.
func (r *tareaRepository) FindByID(id uint) (*Entitys.Tarea, error) {
	var tarea Entitys.Tarea
	err := r.db.Preload("Protocolo").First(&tarea, id).Error
	return &tarea, err
}

// FindAll devuelve una lista de todas las tareas en la base de datos.
func (r *tareaRepository) FindAll() ([]*Entitys.Tarea, error) {
	var tareas []*Entitys.Tarea
	err := r.db.Preload("Protocolo").Find(&tareas).Error
	return tareas, err
}

// Update actualiza una tarea existente en la base de datos.
func (r *tareaRepository) Update(tarea *Entitys.Tarea) error {
	return r.db.Save(tarea).Error
}

// Delete elimina una tarea de la base de datos por su ID.
func (r *tareaRepository) Delete(id uint) error {
	return r.db.Delete(&Entitys.Tarea{}, id).Error
}

// FindByProtocoloID busca todas las tareas asociadas a un protocolo específico por su ID y devuelve una lista de tareas.
func (r *tareaRepository) FindByProtocoloID(protocoloID uint) ([]*Entitys.Tarea, error) {
	var tareas []*Entitys.Tarea
	err := r.db.Where("protocolo_id = ?", protocoloID).Find(&tareas).Error
	return tareas, err
}
