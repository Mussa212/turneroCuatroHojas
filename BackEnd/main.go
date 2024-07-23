package main

import (
	"BackEnd/app"
	"BackEnd/db"
	"BackEnd/jwt"
	log "github.com/sirupsen/logrus"
)

func main() {
	var err error
	jwt.Secreto, err = jwt.GenerateSecretJWT()
	if err != nil {
		log.Debug("Error when creating JWT Token: ", err)
	}
	go db.StartDbEngine()
	go app.StartRoute()
	go jwt.UpdateSecretPeriodically()
}
