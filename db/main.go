package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Database interface {
	Connect(driver, connStr string) error
	Close() error

	AnalyzeSchema(tableName string) (*AnalyzeResult, error)
}

type AnalyzeResult struct {
	IsOptimized bool
}

func (r *AnalyzeResult) String() string {
	return fmt.Sprintf("Is optimized: %v", r.IsOptimized)
}

func NewDatabase(driver string) (Database, error) {
	switch driver {
	case "mysql":
		return NewMySQLDatabase(), nil
	case "postgres":
		return nil, fmt.Errorf("Postgres not implemented yet")
		// return &PostgresDatabase{}, nil
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", driver)
	}
}

func ConnectToDatabase(driver, connectionString string) (*sql.DB, error) {
	db, err := sql.Open(driver, connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %v", err)
	}

	fmt.Println("Connected to database successfully!")
	return db, nil
}
