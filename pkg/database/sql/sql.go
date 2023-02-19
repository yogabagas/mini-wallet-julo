package pql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	DB *sql.DB
}

var (
	SQLClient *MySQL
)

func Connect(user, password, url, schema string) (*MySQL, error) {
	if SQLClient == nil {
		schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, password, url, schema)

		log.Println("SchemaURL", schemaURL)

		db, err := sql.Open("mysql", schemaURL)
		if err != nil {
			log.Panic(err.Error())
			panic(err)
		}

		if err := db.Ping(); err != nil {
			log.Panic(err)
		}
		SQLClient = &MySQL{DB: db}
	}
	return SQLClient, nil
}
