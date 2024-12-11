package main

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

type DBCredential struct {
	HostName     string
	DatabaseName string
	Username     string
	Password     string
	Port         string
}

//TODO: masukkan CAMP_ID kalian dan Credential Database kalian disini
var (
	CAMP_ID = "FS10293520"

	credential = DBCredential{
		
    HostName:  "localhost",
    Port:      "5432",
    Username:      "postgres",
    Password:  "password",
    DatabaseName:  "test_db_adam",


	}
)

func Connection() (db *sql.DB, err error) {
	//setup connection to database postgres
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		credential.HostName, credential.Port, credential.Username, credential.Password, credential.DatabaseName)

	db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        return nil, err
    }

	err = db.Ping()
	if err != nil {
		return
	} else {
		fmt.Println("Successfully connected!")
		err = ioutil.WriteFile("output.txt", []byte(CAMP_ID+" "+fmt.Sprintf("%x", sha256.Sum256([]byte(CAMP_ID)))), 0644)
		return db, err
	}
}

func main() {
	_, err := Connection()

	if err != nil {
		log.Fatal(err)
	}
}
