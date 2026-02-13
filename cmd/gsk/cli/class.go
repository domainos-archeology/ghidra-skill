package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var classCmd = &cobra.Command{
	Use:   "class",
	Short: "Class operations",
}

var listClassesCmd = &cobra.Command{
	Use:   "list",
	Short: "List all classes",
	Long: `List all classes in the program.

Examples:
  # List all classes (up to limit)
  gsk class list

  # Limit results
  gsk class list --limit 100
`,
	Run: func(cmd *cobra.Command, args []string) {
		limit, _ := cmd.Flags().GetInt("limit")
		listClasses(limit)
	},
}

func init() {
	listClassesCmd.Flags().IntP("limit", "l", 1000, "Maximum number of results")

	classCmd.AddCommand(listClassesCmd)
	rootCmd.AddCommand(classCmd)
}

func listClasses(limit int) {
	client := newClient()
	body, err := client.ListClasses(limit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(body))
}
