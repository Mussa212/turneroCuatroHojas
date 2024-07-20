package main

import (
	"BackEnd/app"
	"BackEnd/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
