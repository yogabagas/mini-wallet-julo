package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	db "github.com/yogabagas/jatis-BE/pkg/database/sql"

	"github.com/joho/godotenv"
)

var (
	Db *sql.DB
)

func initModule() {

	//LoadEnv initially load env

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)

	}

	Db = InitDBConn().DB
}

func InitDBConn() *db.MySQL {

	var (
		user, password, host, schema string
	)

	user = os.Getenv("MYSQL_USER")
	password = os.Getenv("MYSQL_PASSWORD")
	host = os.Getenv("MYSQL_ROOT_HOST")
	schema = os.Getenv("MYSQL_DATABASE")

	if user == "" {
		user = "root"
	}
	if password == "" {
		password = ""
	}
	if host == "" {
		host = "localhost"
	}
	if schema == "" {
		schema = "jatisdb"
	}

	mySQL, err := db.Connect(user, password, fmt.Sprintf("%s:3306", host), schema)
	if err != nil {
		return nil
	}
	return mySQL
}
