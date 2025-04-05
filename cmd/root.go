package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Main CLI entry (cobra-based if needed)

func main() {
	fmt.Println("Hello, world!")
	rootCmd := &cobra.Command{
		Use:   "my-app",
		Short: "A simple CLI application",
	}
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Display version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Version: 1.0.0")
		},
	}
	rootCmd.AddCommand(versionCmd)

	rootCmd.Execute()

}
