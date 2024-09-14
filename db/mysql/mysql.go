package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/nhkhang/dba-buddy/ai"
)

type MySQLDatabase struct {
	Conn  *sql.DB
	Agent *ai.OllamaClient

	version string
}

func NewMySQLDatabase(driver string, connStr string, agent *ai.OllamaClient) (*MySQLDatabase, error) {
	db := &MySQLDatabase{
		Agent: agent,
	}
	if err := db.Connect(driver, connStr); err != nil {
		return db, err
	}

	if err := db.Ping(); err != nil {
		return db, err
	}

	return db, nil
}

func (d *MySQLDatabase) Connect(driver, connStr string) error {
	conn, err := sql.Open(driver, connStr)
	if err != nil {
		return err
	}

	d.Conn = conn

	var version string
	err = d.Conn.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		return err
	}
	d.version = version

	return err
}

func (d *MySQLDatabase) Ping() error {
	return d.Conn.Ping()
}

func (d *MySQLDatabase) Close() error {
	return d.Conn.Close()
}

func (d *MySQLDatabase) AnalyzeSchema(tableName string) error {
	rows, err := d.Conn.Query(fmt.Sprintf("DESCRIBE %s", tableName))
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Printf("Columns for table %s:\n", tableName)
	for rows.Next() {
		var field, typeStr, null, key, extra string
		var defaultValue *string
		err := rows.Scan(&field, &typeStr, &null, &key, &defaultValue, &extra)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Field: %s, Type: %s, Null: %s, Key: %s, Default: %v, Extra: %s\n",
			field, typeStr, null, key, defaultValue, extra)
	}

	d.Agent.Analyze("test")

	return nil
}
