package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Se modifica la ruta DB agregando el archivo connection.go
// connection.go se encarga de realizar la conexión a la bbdd
var DSN = "host=localhost user=cam password=2004 dbname=leridapp sslmode=disable" //cadena de conexión
var DBConn *gorm.DB                                                                   //variable DB para llamar GORM

func Connect() error { 
	var err error
	DBConn, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB CONNECTED")
	}
	return err
}