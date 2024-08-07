package controller

import (
	"BackEnd/dto"
	jwtG "BackEnd/jwt"
	service "BackEnd/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetUserById(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var userDto dto.UserDto

	userDto, err := service.UserService.GetUserById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, userDto)
}

func GetUserByEmail(c *gin.Context) {
	email := c.Param("email")

	var userDto dto.UserDto

	userDto, err := service.UserService.GetUserByEmail(email)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, userDto)
}

func GetUsers(c *gin.Context) {
	var usersDto dto.UsersDto
	usersDto, err := service.UserService.GetUsers()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, usersDto)
}

func UserInsert(c *gin.Context) {
	var userDto dto.UserDto
	err := c.BindJSON(&userDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userDto, er := service.UserService.InsertUser(userDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"user_created": "true",
	})
}

func UserAuth(c *gin.Context) {
	var userDto dto.UserDto

	err := c.BindJSON(&userDto)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var autenticado bool
	var tipo int
	var id int
	autenticado, tipo, id = service.UserService.UserAuth(userDto)
	if autenticado == true {
		userDto.Tipo = tipo
		userDto.Id = id
		token, err := jwtG.GenerateUserToken(userDto)
		if err != nil {
			log.Error(err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusAccepted, gin.H{
			"autenticacion": "true",
			"tipo":          tipo,
			"user_id":       id,
			"token":         token,
		})
	} else {
		c.JSON(http.StatusAccepted, gin.H{
			"autenticacion": "false",
			"tipo":          tipo,
			"user_id":       id,
		})
	}

}
