package dto

import "time"

type ReservaDto struct {
	Id     int       `json:"id"`
	UserId int       `json:"user_id,omitempty"`
	Fecha  time.Time `json:"fecha"`
	Tipo   int       `json:"tipo"`
}

type ReservasDto []ReservaDto
