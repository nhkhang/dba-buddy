package cmd

import (
	"fmt"

	"github.com/nhkhang/dba-buddy/db"
)

func handleConnect(driver, host, username, password, dbname string) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, dbname)

	// Create a database instance based on the driver
	var err error
	database, err = db.NewDatabase(driver)
	if err != nil {
		fmt.Println("Error creating database driver:", err)
		return
	}

	// Connect to the database
	if err := database.Connect(driver, connStr); err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	fmt.Println("Connected to the database successfully!")
}

func handleAnalyzeSchema(tableName string) {
	if database == nil {
		fmt.Println("No active database connection")
		return
	}

	// Analyze the schema for the specific table
	result, err := database.AnalyzeSchema(tableName)
	if err != nil {
		fmt.Println("Error analyzing schema for table", tableName, ":", err)
	}

	fmt.Printf("Schema analysis for table %s: %s \n", tableName, result.String())
}
