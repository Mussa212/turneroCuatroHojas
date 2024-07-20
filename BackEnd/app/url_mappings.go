package app

import (
	reservaController "BackEnd/controller/reserva"
	userController "BackEnd/controller/user"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Users Mapping
	router.GET("/reservas", reservaController.GetReservas)
	router.GET("/users", userController.GetUsers)
	router.GET("/users/email", userController.GetUserByEmail)
	router.GET("/users/:id", userController.GetUserById)
	router.POST("/users/auth", userController.UserAuth)
	router.POST("/users", userController.UserInsert)
	router.POST("/reservas", reservaController.ReservaInsert)
	router.POST("/reservas/byfecha", reservaController.GetReservasByFecha)
	router.POST("/reservas/:id", reservaController.GetReservaById)
	router.POST("/reservas/fechauser", reservaController.GetReservasByFechaAndUser)
	router.GET("/reservas/reservauser/:user_id", reservaController.GetReservasByUser)
	router.POST("/reservas/disponibles", reservaController.GetReservasDisponiblesByFecha)
	log.Info("Finishing mappings configurations")
}
