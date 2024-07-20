package dto

type UserDto struct {
	Id        int    `json:"id"`
	FirstName string `json:"name"`
	LastName  string `json:"last_name"`
	UserEmail string `json:"user_email"`
	Password  string `json:"password"`
	Tipo      int    `json:"tipo"`
	DNI       int    `json:"dni"`
	Telefono  string `json:"telefono"`
}

type UsersDto []UserDto
