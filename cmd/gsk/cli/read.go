package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read <address> [length]",
	Short: "Read and display memory bytes in hexadecimal",
	Long: `Read bytes from memory at the specified address and display them
as a hex dump with ASCII representation.

The output format shows:
- Address column
- 16 hex bytes per line (with a gap after 8 bytes)
- ASCII representation (printable characters, '.' for non-printable)

Examples:
  # Read 256 bytes (default) starting at address
  gsk read 0x401234

  # Read specific number of bytes
  gsk read 0x401234 64

  # Read using flag for length
  gsk read 0x401234 --length 512
`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		address := args[0]
		length, _ := cmd.Flags().GetInt("length")

		// If length provided as second positional arg, use that
		if len(args) >= 2 {
			var argLen int
			_, err := fmt.Sscanf(args[1], "%d", &argLen)
			if err == nil && argLen > 0 {
				length = argLen
			}
		}

		readMemory(address, length)
	},
}

func init() {
	readCmd.Flags().IntP("length", "l", 256, "Number of bytes to read (max 65536)")
	rootCmd.AddCommand(readCmd)
}

func readMemory(address string, length int) {
	client := newClient()
	body, err := client.ReadMemory(address, length)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(string(body))
}
