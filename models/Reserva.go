package models

import "gorm.io/gorm"

type Reserva struct {
	gorm.Model
	IdPropiedad   int    `json:"IdPropiedad"`
	FechaInicio   string `json:"FechaInicio"`
	FechaFinal    string `json:"FechaFinal"`
	IdPropietario int    `json:"IdPropietario"`
	IdCliente     int    `json:"IdCliente"`
}
