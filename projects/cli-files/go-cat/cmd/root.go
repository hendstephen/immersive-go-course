package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "go-cat files...",
		Short: "go-cat is basically just cat, but worse",
		Args:  cobra.MinimumNArgs(1),
	}

	cmd.Run = func(cmd *cobra.Command, args []string) {
		for _, path := range args {
			stat, statErr := os.Stat(path)
			if statErr != nil {
				log.Fatalf("File %v does not exist", path)
			}
			if stat.IsDir() {
				log.Fatalf("%v is a directory", path)
			}

			file, _ := os.Open(path)
			defer file.Close()

			// Read file line by line
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}

			if err := scanner.Err(); err != nil {
				log.Fatalf("Error reading file: %v", err)
			}
		}
	}

	return cmd
}

func Execute() {
	NewCmd().Execute()
}
