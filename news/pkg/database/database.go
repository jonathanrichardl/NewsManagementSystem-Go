package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	mysql "news/pkg/database/mysql"
	"strings"
)

type DBInstance struct {
	db *sql.DB
}

type RetrievedData struct {
	Data *sql.Rows
}

func NewDatabase(DatabaseManagementSystem string, Username string, Password string, Address string, DatabaseName string) (*DBInstance, error) {
	var DB *sql.DB
	var err error
	switch strings.ToLower(DatabaseManagementSystem) {
	case "mysql":
		DB, err = mysql.Connect(fmt.Sprintf("%s:%s@tcp(%s)/%s", Username, Password, Address, DatabaseName))
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New(fmt.Sprintf("%s database is not implemented", DatabaseManagementSystem))
	}
	return &DBInstance{DB}, nil
}

func (a *DBInstance) AddData(Query string) error {
	_, err := a.db.Query(Query)
	return err

}

func (a *DBInstance) UpdateData(Query string) error {
	a.db.Query("SET SQL_SAFE_UPDATES = 0;")
	_, err := a.db.Query(Query)
	a.db.Query("SET SQL_SAFE_UPDATES = 1;")
	return err

}

func (a *DBInstance) DeleteData(Query string) error {
	a.db.Query("SET FOREIGN_KEY_CHECKS=0;")
	_, err := a.db.Query(Query)
	a.db.Query("SET FOREIGN_KEY_CHECKS=1;")
	return err
}

func (a *DBInstance) RetrieveData(Query string) (*RetrievedData, error) {
	rows, err := a.db.Query(Query)
	if err != nil {
		return nil, err
	}
	return &RetrievedData{rows}, nil
}

func (a *DBInstance) CheckIfExists(Query string) bool {
	log.Println("Check for duplicates")
	var exists bool
	err := a.db.QueryRow(Query).Scan(&exists)
	if err != nil {
		log.Println("error : ", err.Error())
		return true
	}
	log.Printf("Duplicate returns : %t", exists)
	return exists

}
