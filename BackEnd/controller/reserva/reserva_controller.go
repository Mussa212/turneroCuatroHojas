package controller

import (
	"BackEnd/dto"
	jwtL "BackEnd/jwt"
	service "BackEnd/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func GetReservaById(c *gin.Context) {
	log.Debug("Reserva id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var reservaDto dto.ReservaDto

	reservaDto, err := service.ReservaService.GetReservaById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, reservaDto)
}

func GetReservas(c *gin.Context) {
	var reservasDto dto.ReservasDto
	reservasDto, err := service.ReservaService.GetReservas()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, reservasDto)
}

func ReservaInsert(c *gin.Context) {
	var reservaDto dto.ReservaDto

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token no proporcionado",
		})
		return
	}

	token, err := jwtL.VerifyToken(tokenString)

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token invalido",
		})
		return
	}

	err1 := c.BindJSON(&reservaDto)

	// Error Parsing json param
	if err1 != nil {
		log.Error(err1.Error())
		c.JSON(http.StatusBadRequest, err1.Error())
		return
	}

	reservaDto, er := service.ReservaService.InsertReserva(reservaDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, reservaDto)
}

func GetReservasByUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("user_id"))

	var reservasDto dto.ReservasDto
	reservasDto, err := service.ReservaService.GetReservasByUser(userId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, reservasDto)
}

func GetReservasByFecha(c *gin.Context) {

	var reservaDto dto.ReservaDto

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token no proporcionado",
		})
		return
	}

	token, err := jwtL.VerifyToken(tokenString)

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token invalido",
		})
		return
	}

	err = c.BindJSON(&reservaDto)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var reservasDto dto.ReservasDto

	reservasDto, err = service.ReservaService.GetReservasByFecha(reservaDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, reservasDto)
}

func GetReservasByFechaAndUser(c *gin.Context) {
	var reservaDto dto.ReservaDto

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token no proporcionado",
		})
		return
	}

	token, err := jwtL.VerifyToken(tokenString)

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token invalido",
		})
		return
	}

	err = c.BindJSON(&reservaDto)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var reservasDto dto.ReservasDto

	reservasDto, err = service.ReservaService.GetReservasByFechaAndUser(reservaDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, reservasDto)
}

func GetReservasDisponiblesByFecha(c *gin.Context) {

	var reservaDto dto.ReservaDto

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token no proporcionado",
		})
		return
	}

	token, err := jwtL.VerifyToken(tokenString)

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token invalido",
		})
		return
	}

	err = c.BindJSON(&reservaDto)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var reservasDto dto.ReservasDto

	reservasDto, err = service.ReservaService.GetReservasDisponiblesByFecha(reservaDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, reservasDto)
}
