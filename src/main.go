package main

import (
	"fmt"
	"handlers"
	"log"
	"os"
	"server"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func initDB() *gorm.DB {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS,
		HOST, PORT, DBNAME)
	log.Println(URL)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err)
	}
	log.Println("Database connection established")

	return db
}

func main() {
	db := initDB()
	handlers.InitRepo(db)
	defer db.Close()
	r := server.RegisterRoute()
	r.Run(":8080")
}
