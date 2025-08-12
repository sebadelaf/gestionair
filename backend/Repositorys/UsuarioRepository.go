package Repositorys

import (
	"github.com/sebadelaf/gestionair/Entitys"

	"gorm.io/gorm"
)

type UsuarioRepository interface {
	Create(usuario *Entitys.Usuario) error
	FindByID(id uint) (*Entitys.Usuario, error)
	FindAll() ([]*Entitys.Usuario, error)
	Update(usuario *Entitys.Usuario) error
	Delete(id uint) error
	FindByEmail(email string) (*Entitys.Usuario, error)
}

type usuarioRepository struct {
	db *gorm.DB // Mantiene una referencia a la conexión de la base de datos (GORM).
}

// NewUsuarioRepository crea una nueva instancia de UsuarioRepository con la conexión a la base de datos proporcionada.
func NewUsuarioRepository(db *gorm.DB) UsuarioRepository {
	return &usuarioRepository{db: db}
}

// Create inserta un nuevo usuario en la base de datos.
func (r *usuarioRepository) Create(usuario *Entitys.Usuario) error {
	return r.db.Create(usuario).Error
}

// FindByID busca un usuario por su ID y devuelve un puntero al usuario encontrado o un error si no se encuentra.
func (r *usuarioRepository) FindByID(id uint) (*Entitys.Usuario, error) {
	var usuario Entitys.Usuario
	err := r.db.First(&usuario, id).Error
	return &usuario, err
}

// FindAll devuelve una lista de todos los usuarios en la base de datos.
func (r *usuarioRepository) FindAll() ([]*Entitys.Usuario, error) {
	var usuarios []*Entitys.Usuario
	err := r.db.Find(&usuarios).Error
	return usuarios, err
}

// Update actualiza un usuario existente en la base de datos.
func (r *usuarioRepository) Update(usuario *Entitys.Usuario) error {
	return r.db.Save(usuario).Error
}

// Delete elimina un usuario de la base de datos por su ID.
func (r *usuarioRepository) Delete(id uint) error {
	return r.db.Delete(&Entitys.Usuario{}, id).Error
}

// FindByEmail busca un usuario por su email y devuelve un puntero al usuario encontrado o un error si no se encuentra.
func (r *usuarioRepository) FindByEmail(email string) (*Entitys.Usuario, error) {
	var usuario Entitys.Usuario
	err := r.db.Where("email = ?", email).First(&usuario).Error
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}
