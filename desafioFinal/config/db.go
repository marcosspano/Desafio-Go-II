package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var URLDbConnection = ""

func LoadConfig() {

	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	URLDbConnection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_URL"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME"))

}

func ConnectDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", URLDbConnection)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
