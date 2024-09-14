package db

import (
	"database/sql"
	"fmt"
)

type MySQLDatabase struct {
	Conn *sql.DB
}

func NewMySQLDatabase() *MySQLDatabase {
	return &MySQLDatabase{}
}

func (db *MySQLDatabase) Connect(driver, connStr string) error {
	conn, err := sql.Open(driver, connStr)
	if err != nil {
		return err
	}

	db.Conn = conn

	return err
}

func (db *MySQLDatabase) Close() error {
	return db.Conn.Close()
}

func (db *MySQLDatabase) AnalyzeSchema(tableName string) (*AnalyzeResult, error) {
	return nil, fmt.Errorf("Not implemented yet")
}
