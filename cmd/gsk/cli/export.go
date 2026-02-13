package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export operations (entry points)",
}

var listExportsCmd = &cobra.Command{
	Use:   "list",
	Short: "List exported entry points",
	Long: `List exported entry points in the program.

Examples:
  # List all exports (up to limit)
  gsk export list

  # Filter by name
  gsk export list --filter main

  # Limit results
  gsk export list --limit 100
`,
	Run: func(cmd *cobra.Command, args []string) {
		filter, _ := cmd.Flags().GetString("filter")
		limit, _ := cmd.Flags().GetInt("limit")
		listExports(filter, limit)
	},
}

func init() {
	listExportsCmd.Flags().StringP("filter", "f", "", "Filter exports by name")
	listExportsCmd.Flags().IntP("limit", "l", 1000, "Maximum number of results")

	exportCmd.AddCommand(listExportsCmd)
	rootCmd.AddCommand(exportCmd)
}

func listExports(filter string, limit int) {
	client := newClient()
	body, err := client.ListExports(filter, limit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(body))
}
