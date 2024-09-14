package cmd

import (
	"fmt"

	"github.com/nhkhang/dba-buddy/db"
	"github.com/spf13/cobra"
)

var scriptModeCommands = []*cobra.Command{
	analyzeSchemaCmd,
	// Fill other script mode commands here...
}

func addDBConnectionFlags(cmd *cobra.Command) {
	dbConnectionFlags := []Flag{
		{
			Name:       FlagNameDriver,
			Usage:      "Database driver (mysql, postgres, etc.)",
			IsRequired: true,
		},
		{
			Name:       FlagNameHost,
			Usage:      "Database host",
			IsRequired: true,
		},
		{
			Name:       FlagNameUsername,
			Usage:      "Database username",
			IsRequired: true,
		},
		{
			Name:       FlagNamePassword,
			Usage:      "Database password",
			IsRequired: true,
		},
		{
			Name:       FlagNameDatabase,
			Usage:      "Database name",
			IsRequired: true,
		},
	}

	for _, flag := range dbConnectionFlags {
		cmd.Flags().String(flag.Name, "", flag.Usage)
		if flag.IsRequired {
			cmd.MarkFlagRequired(flag.Name)
		}
	}
}

func GetConnStr(cmd *cobra.Command) string {
	host, _ := cmd.Flags().GetString("host")
	username, _ := cmd.Flags().GetString("username")
	password, _ := cmd.Flags().GetString("password")
	dbname, _ := cmd.Flags().GetString("dbname")

	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, dbname)
}

var analyzeSchemaCmd = &cobra.Command{
	Use:   "analyze-schema [table-name]",
	Short: "Analyze the schema of a table",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		driver, _ := cmd.Flags().GetString("driver")
		connStr := getConnStr(cmd)

		database, err := db.NewDatabase(driver)
		if err != nil {
			fmt.Println(err)
			return
		}

		if err := database.Connect(driver, connStr); err != nil {
			fmt.Println("Error connecting to the database:", err)
			return
		}
		fmt.Println("Connected to the database successfully!")

		// Analyze schema
		handleAnalyzeSchema(args[0])

		// Close connection after the operation
		if err := database.Close(); err != nil {
			fmt.Println("Error closing the database connection:", err)
		} else {
			fmt.Println("Disconnected from the database.")
		}
	},
}
