package dataBase

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "library"
)

func InitDataBase(){
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	if err!= nil{
		log.Fatal(err)
	}
	defer db.Close()


}