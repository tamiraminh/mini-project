package database


import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:13519129@tcp(127.0.0.1:3306)/bootcampcrud")
	if err != nil {
		return nil, err
	}

	return db, nil
}