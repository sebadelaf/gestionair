package config

import (
	"fmt"
	"github.com/sebadelaf/gestionair/Entitys"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

// DB es una variable exportada que mantendrá la instancia de la conexión a la base de datos.
// Otras partes de la aplicación (como los repositorios) la usarán para interactuar con la BD.
var DB *gorm.DB

// ConnectDatabase es la función que inicializa la conexión con GORM.
func ConnectDatabase() {
	// 1. Carga la URL de la base de datos desde el archivo .env
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		panic("La variable de entorno DB_URL no está definida.")
	}

	// 2. Abre la conexión con la base de datos usando GORM
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// Si la conexión falla, la aplicación no puede continuar.
		// `panic` detendrá la ejecución y mostrará el error.
		panic(fmt.Sprintf("Fallo al conectar a la base de datos: %v", err))
	}

	fmt.Println("Conexión a la base de datos exitosa.")

	// 3. Auto-Migración: GORM crea o actualiza las tablas en la BD
	// para que coincidan con los structs de tus entidades.
	// ¡Debes listar aquí TODAS tus entidades!
	err = database.AutoMigrate(
		&Entitys.Usuario{},
		&Entitys.Cliente{},
		&Entitys.Equipo{},
		&Entitys.Protocolo{},
		&Entitys.Tarea{},
		&Entitys.OrdenDeTrabajo{},
		&Entitys.OrdenDeServicio{},
	)
	if err != nil {
		panic(fmt.Sprintf("Fallo al migrar la base de datos: %v", err))
	}
	fmt.Println("Migración de la base de datos exitosa.")

	// 4. Asigna la instancia de la base de datos a la variable global DB
	// para que pueda ser accedida desde otros paquetes.
	DB = database
}
