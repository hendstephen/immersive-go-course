// cmd/root.go
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "go-ls [path]",
		Short: "go-ls is basically just ls, but worse",
		Args:  cobra.RangeArgs(0, 1),
	}

	mFlag := cmd.Flags().BoolP("m", "m", false, "fill width with a comma separated list of entries")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		// Default to working dir
		dir := "."

		// If path is provided, use that
		if len(args) > 0 {
			dir = args[0]
		}

		// Check if path exists
		stat, err := os.Stat(dir)
		if err != nil {
			return fmt.Errorf("file or directory %v does not exist", dir)
		}

		// If it's a file, print the file
		if !stat.IsDir() {
			fmt.Fprintln(cmd.OutOrStdout(), stat.Name())
			return nil
		}

		// If a dir, print files in the dir
		files, err := os.ReadDir(dir)
		if err != nil {
			return fmt.Errorf("error reading directory: %v", err)
		}

		if *mFlag {
			filenames := make([]string, 0, len(files))
			for _, file := range files {
				filenames = append(filenames, file.Name())
			}

			fmt.Println(strings.Join(filenames, ", "))
		} else {
			for _, file := range files {
				fmt.Fprintln(cmd.OutOrStdout(), file.Name())
			}
		}

		return nil
	}

	return cmd
}

func Execute() {
	cmd := NewCmd()

	// Cobra will handle printing the error
	cmd.Execute()
}
