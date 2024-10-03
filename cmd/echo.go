package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var echoCmd = &cobra.Command{
	Use: "echo [text]",
	Short: Echo the input text,
	Args: cobra.MinimumArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println(args[0])
    },
}

funct init() {
    RootCmd.AddCommand(echoCmd)
}