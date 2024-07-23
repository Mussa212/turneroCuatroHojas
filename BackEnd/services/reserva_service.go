package services

import (
	e "BackEnd/Utils"
	reservaClient "BackEnd/clients/reserva"
	"BackEnd/dto"
	"BackEnd/model"
	"time"
)

type reservaService struct{}

type reservaServiceInterface interface {
	GetReservaById(id int) (dto.ReservaDto, e.ApiError)
	GetReservas() (dto.ReservasDto, e.ApiError)
	InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError)
	GetReservasByUser(userId int) (dto.ReservasDto, e.ApiError)
	GetReservasByFecha(reservaDto dto.ReservaDto) (dto.ReservasDto, e.ApiError)
	GetReservasByFechaAndUser(reservaDto dto.ReservaDto) (dto.ReservasDto, e.ApiError)
	GetReservasDisponiblesByFecha(reservaDto dto.ReservaDto) (dto.ReservasDto, e.ApiError)
}

var (
	ReservaService reservaServiceInterface
)

func init() {
	ReservaService = &reservaService{}
}

func (s *reservaService) GetReservaById(id int) (dto.ReservaDto, e.ApiError) {

	var reserva model.Reserva = reservaClient.ReservaClient.GetReservaById(id)
	var reservaDto dto.ReservaDto

	if reserva.ID == 0 {
		return reservaDto, e.NewBadRequestApiError("reserva not found")
	}

	reservaDto.Fecha = reserva.Fecha
	reservaDto.UserId = reserva.UserId
	reservaDto.Id = reserva.ID
	reserva.Tipo = reserva.Tipo

	return reservaDto, nil
}

func (s *reservaService) GetReservas() (dto.ReservasDto, e.ApiError) {

	var reservas model.Reservas = reservaClient.ReservaClient.GetReservas()
	var reservasDto dto.ReservasDto

	for _, reserva := range reservas {
		var reservaDto dto.ReservaDto

		reservaDto.Fecha = reserva.Fecha
		reservaDto.UserId = reserva.UserId
		reservaDto.Id = reserva.ID
		reservaDto.Tipo = reserva.Tipo
		reservasDto = append(reservasDto, reservaDto)
	}

	return reservasDto, nil
}

func (s *reservaService) InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError) {

	var reserva model.Reserva

	reserva.Fecha = reservaDto.Fecha
	reserva.UserId = reservaDto.UserId
	reserva.Tipo = reservaDto.Tipo

	reserva = reservaClient.ReservaClient.InsertReserva(reserva)

	reservaDto.Id = reserva.ID

	return reservaDto, nil
}

func (s *reservaService) GetReservasByUser(userId int) (dto.ReservasDto, e.ApiError) {

	var reservas model.Reservas = reservaClient.ReservaClient.GetReservasByUser(userId)
	var reservasDto dto.ReservasDto

	for _, reserva := range reservas {
		var reservaDto dto.ReservaDto

		reservaDto.Fecha = reserva.Fecha
		reservaDto.UserId = reserva.UserId
		reservaDto.Id = reserva.ID
		reservaDto.Tipo = reserva.Tipo
		reservasDto = append(reservasDto, reservaDto)
	}

	return reservasDto, nil
}

func (s *reservaService) GetReservasByFecha(reservaDto dto.ReservaDto) (dto.ReservasDto, e.ApiError) {

	var reserva model.Reserva

	reserva.Fecha = reservaDto.Fecha

	var reservas model.Reservas = reservaClient.ReservaClient.GetReservasByFecha(reserva)
	var reservasDto dto.ReservasDto

	for _, reserva = range reservas {
		var reservaDto dto.ReservaDto

		reservaDto.Fecha = reservaDto.Fecha
		reservaDto.UserId = reserva.UserId
		reservaDto.Id = reserva.ID
		reservaDto.Tipo = reserva.Tipo
		reservasDto = append(reservasDto, reservaDto)
	}

	return reservasDto, nil
}

func (s *reservaService) GetReservasByFechaAndUser(reservaDto dto.ReservaDto) (dto.ReservasDto, e.ApiError) {
	var reserva model.Reserva

	reserva.Fecha = reservaDto.Fecha
	reserva.UserId = reservaDto.UserId

	var reservas model.Reservas = reservaClient.ReservaClient.GetReservasByFechaAndUser(reserva)
	var reservasDto dto.ReservasDto

	for _, reserva = range reservas {
		var reservaDto dto.ReservaDto

		reservaDto.Fecha = reserva.Fecha
		reservaDto.UserId = reserva.UserId
		reservaDto.Id = reserva.ID
		reservaDto.Tipo = reserva.Tipo
		reservasDto = append(reservasDto, reservaDto)
	}

	return reservasDto, nil
}

func (s *reservaService) GetReservasDisponiblesByFecha(reservaDto dto.ReservaDto) (dto.ReservasDto, e.ApiError) {

	var reserva model.Reserva

	reserva.Fecha = reservaDto.Fecha

	var reservas model.Reservas = reservaClient.ReservaClient.GetReservasByFecha(reserva)
	var reservasDto dto.ReservasDto

	recorrido := time.Date(reserva.Fecha.Year(), reserva.Fecha.Month(), reserva.Fecha.Day(), 9, 0, 0, 0, reserva.Fecha.Location())

	for i := 0; i < 11; i++ {
		var reservaRecorridoDto dto.ReservaDto
		if recorrido != reservas[0].Fecha && recorrido.Hour() <= 19 {
			reservaRecorridoDto.Fecha = recorrido
			reservasDto = append(reservasDto, reservaRecorridoDto)
		}
		recorrido = recorrido.Add(1 * time.Hour)
	}

	return reservasDto, nil
}
