package Entitys

type Protocolos struct {
	IdProtocolo int    `json:"id_protocolo" db:"id_protocolo"`
	Nombre      string `json:"nombre" db:"nombre"`
}
