package Entitys

type Equipos struct {
	ID                int    `json:"id" db:"id"`
	ID_cliente        int    `json:"id_cliente" db:"id_cliente"`
	Nombre            string `json:"nombre" db:"nombre"`
	Marca             string `json:"marca" db:"marca"`
	Modelo            string `json:"modelo" db:"modelo"`
	NumeroSerie       string `json:"numero_serie" db:"numero_serie"`
	Tipo              string `json:"tipo" db:"tipo"`
	Capacidad         string `json:"capacidad" db:"capacidad"`
	Refrigeramte      string `json:"refrigerante" db:"refrigerante"`
	Tecnologia        string `json:"tecnologia" db:"tecnologia"`
	Voltaje           string `json:"voltaje" db:"voltaje"`
	Fecha_instalacion string `json:"fecha_instalacion" db:"fecha_instalacion"`
	Ubicacion         string `json:"ubicacion" db:"ubicacion"`
	ID_protocolos     int    `json:"id_protocolo" db:"id_protocolo"`
}
