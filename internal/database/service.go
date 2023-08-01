package database


import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "<username>:<password>@tcp(127.0.0.1:<mysqlserverport>)/<databasename>")
	if err != nil {
		return nil, err
	}

	return db, nil
}