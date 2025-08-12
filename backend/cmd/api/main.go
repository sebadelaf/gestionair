package main

import (
	"log"
	// Paquetes de terceros
	//"github.com/gin-contrib/cors"
	//"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sebadelaf/gestionair/Repositorys"
	// Paquetes propios
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

	// 3. Repositorios

	// Inicializa el repositorio de clientes.
	clienteRepo := Repositorys.NewClienteRepository(config.DB)
	if clienteRepo != nil {
		log.Println("¡El repositorio de clientes fue inicializado correctamente!")
	} else {
		log.Println("Error al inicializar el repositorio de clientes.")
	}
	// Inicializa el repositorio de equipos.
	equipoRepo := Repositorys.NewEquipoRepository(config.DB)
	if equipoRepo != nil {
		log.Println("¡El repositorio de equipos fue inicializado correctamente!")
	} else {
		log.Println("Error al inicializar el repositorio de equipos.")
	}
	// Inicializa el repositorio de protocolos.
	protocoloRepo := Repositorys.NewProtocoloRepository(config.DB)
	if protocoloRepo != nil {
		log.Println("¡El repositorio de protocolos fue inicializado correctamente!")
	} else {
		log.Println("Error al inicializar el repositorio de protocolos.")
	}
	//protoclo de tareas
	tareaRepo := Repositorys.NewTareaRepository(config.DB)
	if tareaRepo != nil {
		log.Println("¡El repositorio de tareas fue inicializado correctamente!")
	} else {
		log.Println("Error al inicializar el repositorio de tareas.")
	}
	// Inicializa el repositorio de usuarios.
	usuarioRepo := Repositorys.NewUsuarioRepository(config.DB)
	if usuarioRepo != nil {
		log.Println("¡El repositorio de usuarios fue inicializado correctamente!")
	} else {
		log.Println("Error al inicializar el repositorio de usuarios.")
	}
	// repositorio de orden de trabajo
	ordenTrabajoRepo := Repositorys.NewOrdenDeTrabajoRepository(config.DB)
	if ordenTrabajoRepo != nil {
		log.Println("¡El repositorio de orden de trabajo fue inicializado correctamente!")
	} else {
		log.Println("Error al inicializar el repositorio de orden de trabajo.")
	}
	// repositorio de orden de servicio
	ordenServicioRepo := Repositorys.NewOrdenDeServicioRepository(config.DB)
	if ordenServicioRepo != nil {
		log.Println("¡El repositorio de orden de servicio fue inicializado correctamente!")
	} else {
		log.Println("Error al inicializar el repositorio de orden de servicio.")
	}
}
