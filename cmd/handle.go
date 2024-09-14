package cmd

import (
	"fmt"

	"github.com/nhkhang/dba-buddy/ai"
	"github.com/nhkhang/dba-buddy/db"
)

type ConnectionConfig struct {
	Driver      string
	DBConnStr   string
	AIAgentHost string
}

func handleConnect(driver, connStr string) error {
	// Map AI Agent
	agent, err := ai.NewOllamaClient()
	if err != nil {
		return err
	}

	database, err = db.NewDatabase(driver, connStr, agent)
	if err != nil {
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

	err := database.AnalyzeSchema(tableName)
	if err != nil {
		fmt.Println("Error analyzing schema for table", tableName, ":", err)
	}
}
