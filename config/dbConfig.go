package config

import (
	"fmt"

	"crud.com/models/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

//SetupDatabaseConnection is create a connection to database when server boot up
func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env")
	}

	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "root"
	dbname := "golang"

	//psqlInfo koneks
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		println(err.Error())
		panic("Cannot connect to database!")
	}
	println("Db is connected!")
	db.AutoMigrate(&entity.Barang{})
	db.AutoMigrate(&entity.Kategori{})
	return db
}

//CloseDatabaseConnection will close connection to database
// func CloseDatabaseConnection(db *gorm.DB) {
// 	dbSQL, err := db.DB()
// 	if err != nil {
// 		panic("Failed when close a connection from database")
// 	}
// 	// defer dbSql.Close()
// 	dbSQL.Close()
// }
