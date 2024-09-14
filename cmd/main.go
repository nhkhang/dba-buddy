package cmd

import (
	"fmt"
	"os"

	"github.com/nhkhang/dba-buddy/db"
	"github.com/spf13/cobra"
)

var database db.Database

var RootCmd = &cobra.Command{
	Use:   "dbbuddy",
	Short: "A Database Administrator Buddy",
}

// Execute is the entry point of the application
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getConnStr(cmd *cobra.Command) string {
	host, _ := cmd.Flags().GetString(FlagNameHost)
	username, _ := cmd.Flags().GetString(FlagNameUsername)
	password, _ := cmd.Flags().GetString(FlagNamePassword)
	dbname, _ := cmd.Flags().GetString(FlagNameDatabase)

	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, dbname)
}

func init() {
	commands := []*cobra.Command{
		startCmd,
	}
	commands = append(commands, scriptModeCommands...)

	// Script mode
	for _, cmd := range commands {
		RootCmd.AddCommand(cmd)
		addDBConnectionFlags(cmd)
	}
}
