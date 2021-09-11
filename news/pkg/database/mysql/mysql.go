package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(Details string) (*sql.DB, error) {
	db, err := sql.Open("mysql", Details)
	if err != nil {
		return nil, err
	}
	return db, nil
}
