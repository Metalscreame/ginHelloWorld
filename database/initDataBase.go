package dataBase

import (
	"database/sql"
	"fmt"
	"time"
	"log"
	_"github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "root"
	DB_NAME     = "postgres"
)

/*
CREATE TABLE users( id serial NOT NULL, email character varying(100) NOT NULL,username character varying(100) NOT NULL,password character varying(100) NOT NULL,
        firstname character varying(500) NOT NULL,
	lastname character varying(500) NOT NULL,
        Created date,
        CONSTRAINT userinfo_pkey PRIMARY KEY (id))
    WITH (OIDS=FALSE);
 */

func InitDataBase() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	insertUser(db,"roma@roma.com","romka","passwordddd","Roma","Kosyiy","2012-12-09")
	updateUserName(db,"newUserName",1)
	selectAllFromUsers(db)
	deleteUser(db,2)
}




func updateUserName(db *sql.DB, username string,id int){
	fmt.Println("# Updating")
	stmt, err := db.Prepare("update users set username=$1 where id=$2")
	checkErr(err)
	res, err := stmt.Exec(username, id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect, "user have been updated")
}


func insertUser(db *sql.DB, email,username,password,firstname,lastname,created string){
	var lastInsertId int
	fmt.Println("Inserting new user")
	err := db.QueryRow("INSERT INTO users(email,username,password,firstname,lastname,Created) VALUES($1,$2,$3,$4,$5,$6) returning id;", email,
		username,password,firstname,lastname,created).Scan(&lastInsertId)
	checkErr(err)
	fmt.Println("last inserted id =", lastInsertId)
}

func deleteUser(db * sql.DB, id int){
	fmt.Println("Deleting")
	stmt, err := db.Prepare("delete from users where id=$1")
	checkErr(err)
	res, err := stmt.Exec(id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect, "deleted")
}

func selectAllFromUsers(db * sql.DB){
	fmt.Println("# Querying")
	rows, err := db.Query("SELECT * FROM users ORDER BY id")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username,email,password string
		var firstname,lastname string
		var created time.Time
		err = rows.Scan(&uid, &email, &username, &password,&firstname,&lastname,&created)///email,username,password,firstname,lastname,Created
		checkErr(err)
		fmt.Println("id | username | department | created ")
		fmt.Printf("%3v | %8v | %6v | %6v %3v | %8v | %6v\n", uid,email, username,password, firstname,lastname, created)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
