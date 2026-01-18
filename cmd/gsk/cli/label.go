package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var labelCmd = &cobra.Command{
	Use:   "label",
	Short: "Label operations (add/remove symbols at addresses)",
}

var listLabelsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all labels or labels at a specific address",
	Long: `List all labels in the program, or filter by address.

Examples:
  # List all labels (up to limit)
  gsk label list

  # List labels at a specific address
  gsk label list --address 0x401234

  # Limit results
  gsk label list --limit 100
`,
	Run: func(cmd *cobra.Command, args []string) {
		address, _ := cmd.Flags().GetString("address")
		limit, _ := cmd.Flags().GetInt("limit")
		listLabels(address, limit)
	},
}

var addLabelCmd = &cobra.Command{
	Use:   "add <address> <name>",
	Short: "Add a label at an address",
	Long: `Create a label at the specified address.

Labels can be global or local (function-scoped). By default, labels
are created in the global namespace. Use --local to create a label
scoped to the containing function.

Examples:
  # Create a global label
  gsk label add 0x401234 loop_start

  # Create a local label (scoped to containing function)
  gsk label add 0x401234 inner_loop --local

  # Explicitly create a global label
  gsk label add 0x401234 global_marker --global
`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		local, _ := cmd.Flags().GetBool("local")
		scope := ""
		if local {
			scope = "local"
		}
		addLabel(args[0], args[1], scope)
	},
}

var deleteLabelCmd = &cobra.Command{
	Use:   "delete <address> <name>",
	Short: "Delete a label at an address",
	Long: `Remove a label from the specified address.

Examples:
  # Delete a label
  gsk label delete 0x401234 loop_start
`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		deleteLabel(args[0], args[1])
	},
}

func init() {
	listLabelsCmd.Flags().StringP("address", "a", "", "Filter by address")
	listLabelsCmd.Flags().IntP("limit", "l", 1000, "Maximum number of results")

	addLabelCmd.Flags().BoolP("local", "L", false, "Create label in function-local scope")
	addLabelCmd.Flags().BoolP("global", "g", false, "Create label in global scope (default)")

	labelCmd.AddCommand(listLabelsCmd)
	labelCmd.AddCommand(addLabelCmd)
	labelCmd.AddCommand(deleteLabelCmd)
	rootCmd.AddCommand(labelCmd)
}

func listLabels(address string, limit int) {
	client := newClient()
	body, err := client.ListLabels(address, limit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(body))
}

func addLabel(address, name, scope string) {
	client := newClient()
	body, err := client.SetLabel(address, name, scope)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(body))
}

func deleteLabel(address, name string) {
	client := newClient()
	body, err := client.DeleteLabel(address, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(body))
}
