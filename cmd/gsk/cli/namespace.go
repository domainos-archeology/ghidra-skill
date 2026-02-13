package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var namespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "Namespace operations",
}

var listNamespacesCmd = &cobra.Command{
	Use:   "list",
	Short: "List all namespaces",
	Long: `List all namespaces in the program.

Examples:
  # List all namespaces (up to limit)
  gsk namespace list

  # Limit results
  gsk namespace list --limit 100
`,
	Run: func(cmd *cobra.Command, args []string) {
		limit, _ := cmd.Flags().GetInt("limit")
		listNamespaces(limit)
	},
}

func init() {
	listNamespacesCmd.Flags().IntP("limit", "l", 1000, "Maximum number of results")

	namespaceCmd.AddCommand(listNamespacesCmd)
	rootCmd.AddCommand(namespaceCmd)
}

func listNamespaces(limit int) {
	client := newClient()
	body, err := client.ListNamespaces(limit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(body))
}
