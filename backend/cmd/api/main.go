package main

import (
	"log"
	// Paquetes de terceros
	//"github.com/gin-contrib/cors"
	//"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// Paquetes propios
	//"github.com/sebadelaf/gestionair/Entitys"
	"github.com/sebadelaf/gestionair/config"
)

func main() {
	// 1. --- Carga de Configuración ---
	// Carga las variables de entorno desde el archivo .env en la raíz del proyecto.
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error: No se pudo cargar el archivo .env: %v", err)
	}

	// 2. --- Conexión a la Base de Datos ---
	// Inicializa la conexión con GORM y ejecuta AutoMigrate para crear/actualizar las tablas.
	// La variable config.DB quedará disponible para ser usada por los repositorios.
	config.ConnectDatabase()

}
