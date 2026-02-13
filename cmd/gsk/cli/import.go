package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import operations (external/library symbols)",
}

var listImportsCmd = &cobra.Command{
	Use:   "list",
	Short: "List imported symbols",
	Long: `List imported (external) symbols in the program.

Examples:
  # List all imports (up to limit)
  gsk import list

  # Filter by name
  gsk import list --filter printf

  # Limit results
  gsk import list --limit 100
`,
	Run: func(cmd *cobra.Command, args []string) {
		filter, _ := cmd.Flags().GetString("filter")
		limit, _ := cmd.Flags().GetInt("limit")
		listImports(filter, limit)
	},
}

func init() {
	listImportsCmd.Flags().StringP("filter", "f", "", "Filter imports by name")
	listImportsCmd.Flags().IntP("limit", "l", 1000, "Maximum number of results")

	importCmd.AddCommand(listImportsCmd)
	rootCmd.AddCommand(importCmd)
}

func listImports(filter string, limit int) {
	client := newClient()
	body, err := client.ListImports(filter, limit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(body))
}
