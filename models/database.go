package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDb() *sql.DB {
	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")

	connection_string := fmt.Sprintf("postgresql://"+"%s"+":"+"%s"+"@"+"%s"+":"+"%v"+"/"+"%s", db_user, db_password, db_host, db_port, db_name)
	db, err := sql.Open("pgx", connection_string)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db

}
