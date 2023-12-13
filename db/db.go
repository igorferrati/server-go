package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnPostgresDb() *sql.DB {
	conn := "host=localhost user=postgres password=i3am18hc dbname=loja_go sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}
	return db
}
