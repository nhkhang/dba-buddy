package db

import (
	"database/sql"
	"fmt"
	"log"
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

func (db *MySQLDatabase) Ping() error {
	return db.Conn.Ping()
}

func (db *MySQLDatabase) Close() error {
	return db.Conn.Close()
}

func (db *MySQLDatabase) AnalyzeSchema(tableName string) (*AnalyzeResult, error) {
	query := fmt.Sprintf("SHOW CREATE TABLE %s", tableName)
	rows, err := db.Conn.Query(query)
	if err != nil {
		log.Fatal("Error executing query: ", err)
	}
	defer rows.Close()

	// Iterate over the result set
	for rows.Next() {
		var tableName, createTableSQL string
		if err := rows.Scan(&tableName, &createTableSQL); err != nil {
			log.Fatal("Error scanning row: ", err)
		}
		fmt.Println("Table Name:", tableName)
		fmt.Println("Create Table SQL:")
		fmt.Println(createTableSQL)
	}

	return nil, fmt.Errorf("Not implemented yet")
}
