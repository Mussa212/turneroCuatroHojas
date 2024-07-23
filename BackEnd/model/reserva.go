package model

import "time"

type Reserva struct {
	ID     int       `gorm:"primaryKey"`
	Fecha  time.Time `gorm:"type:DATETIME"`
	UserId int       `gorm:"foreignKey"`
	Tipo   int       `gorm:"int"`
}

type Reservas []Reserva
