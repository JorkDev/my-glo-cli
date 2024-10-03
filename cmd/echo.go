package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var echoCmd = &cobra.Command{
    Use:   "echo [text]",
    Short: "Echo the input text",
    Args:  cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println(args[0])
    },
}

func init() {
    RootCmd.AddCommand(echoCmd)
}
