package db

import (
	reservaClient "BackEnd/clients/reserva"
	userClient "BackEnd/clients/user"
	"BackEnd/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var (
	db  *gorm.DB
	err error
)

func insertInitialData() {
	// Insert users
	user := model.User{
		FirstName: "Admin",
		LastName:  "Admin",
		Email:     "admin@admin.com",
		Password:  "password123",
		Tipo:      1,
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Error al hashear la password:", err.Error())
	}
	user.Password = string(hashedPassword)
	if err := db.Create(&user).Error; err != nil {
		log.Error("Failed to insert user:", err.Error())
	}

	log.Info("Initial values inserted")
}

func init() {
	// DB Connections Paramters
	DBName := "pruebaHash"
	DBUser := "root"
	DBPass := "arquisoft1"
	//DBPass := os.Getenv("MVC_DB_PASS")
	DBHost := "localhost"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	userClient.Db = db
	reservaClient.Db = db

}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Reservas{})

	insertInitialData()

	log.Info("Finishing Migration Database Tables")
}
