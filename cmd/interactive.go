package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// startCmd starts the interactive shell
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start dbbuddy and enter interactive mode",
	Run: func(cmd *cobra.Command, args []string) {
		driver, _ := cmd.Flags().GetString("driver")
		connStr := getConnStr(cmd)

		handleConnect(driver, connStr)

		StartInteractiveShell()
	},
}

func StartInteractiveShell() {
	fmt.Println("Starting dba-buddy interactive mode. Type 'exit' to quit.")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ") // Prompt symbol
		scanner.Scan()
		input := scanner.Text()

		// Handle user input and split it into command parts
		command := strings.Fields(input)
		if len(command) == 0 {
			continue
		}

		// Exit command to break the loop
		if command[0] == "exit" {
			fmt.Println("Exiting dbbuddy...")
			break
		}

		// Process commands
		switch command[0] {
		case "analyze-schema":
			if len(command) != 2 {
				fmt.Println("Usage: analyze-schema [table-name]")
				continue
			}
			handleAnalyzeSchema(command[1])
		default:
			fmt.Println("Unknown command:", command[0])
		}
	}
}
