package cmd

import (
	"fmt"

	"github.com/nhkhang/dba-buddy/db"
)

func handleConnect(driver, connStr string) error {
	var err error
	database, err = db.NewDatabase(driver)
	if err != nil {
		return err
	}

	if err := database.Connect(driver, connStr); err != nil {
		return err
	}
	defer func() {
		if err := database.Close(); err != nil {
			fmt.Println("Error closing database connection:", err)
		}
	}()

	if err = database.Ping(); err != nil {
		return err
	}

	fmt.Println("Connected to the database successfully!")

	return nil
}

func handleAnalyzeSchema(tableName string) {
	if database == nil {
		fmt.Println("You need to connect to the database first")
		return
	}

	result, err := database.AnalyzeSchema(tableName)
	if err != nil {
		fmt.Println("Error analyzing schema for table", tableName, ":", err)
	}

	fmt.Printf("Schema analysis for table %s: %s \n", tableName, result.String())
}
