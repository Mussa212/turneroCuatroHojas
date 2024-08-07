package services

import (
	e "BackEnd/Utils"
	userClient "BackEnd/clients/user"
	"BackEnd/dto"
	"BackEnd/model"
	"golang.org/x/crypto/bcrypt"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (dto.UserDto, e.ApiError)
	GetUsers() (dto.UsersDto, e.ApiError)
	InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError)
	GetUserByEmail(email string) (dto.UserDto, e.ApiError)
	UserAuth(userDto dto.UserDto) (bool, int, int)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (dto.UserDto, e.ApiError) {

	var user model.User = userClient.GetUserById(id)
	var userDto dto.UserDto

	if user.ID == 0 {
		return userDto, e.NewBadRequestApiError("user not found")
	}

	userDto.FirstName = user.FirstName
	userDto.LastName = user.LastName
	userDto.UserEmail = user.Email
	userDto.Tipo = user.Tipo
	userDto.Id = user.ID
	userDto.DNI = user.DNI
	userDto.Telefono = user.Telefono

	return userDto, nil
}

func (s *userService) GetUserByEmail(email string) (dto.UserDto, e.ApiError) {
	var user model.User = userClient.GetUserByEmail(email)
	var userDto dto.UserDto

	if user.Email == "" {
		return userDto, e.NewBadRequestApiError("user not found")
	}

	userDto.FirstName = user.FirstName
	userDto.LastName = user.LastName
	userDto.Tipo = user.Tipo
	userDto.Id = user.ID
	userDto.UserEmail = user.Email
	userDto.DNI = user.DNI
	userDto.Telefono = user.Telefono
	return userDto, nil
}

func (s *userService) GetUsers() (dto.UsersDto, e.ApiError) {

	var users model.Users = userClient.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.FirstName = user.FirstName
		userDto.LastName = user.LastName
		userDto.Id = user.ID
		userDto.Tipo = user.Tipo
		userDto.DNI = user.DNI
		userDto.Telefono = user.Telefono

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)

	if err != nil {
		// Handle error
	}
	var user model.User

	user.FirstName = userDto.FirstName
	user.LastName = userDto.LastName
	user.Password = string(hashedPassword)
	user.Email = userDto.UserEmail
	user.Tipo = userDto.Tipo
	user.DNI = userDto.DNI
	user.Telefono = userDto.Telefono

	user = userClient.InsertUser(user)

	userDto.Id = user.ID

	return userDto, nil
}

func (s *userService) UserAuth(userDto dto.UserDto) (bool, int, int) {

	user := userClient.GetUserByEmail(userDto.UserEmail)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDto.Password))
	if err != nil {
		return false, -1, -1
	}

	return true, user.Tipo, user.ID
}
