package cmd

import (
    "github.com/spf13/cobra"
    "fmt"
)

var RootCmd = &cobra.Command{
    Use:   "mycli",
    Short: "MyCLI is a simple CLI",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Welcome to MyCLI")
    },
}

func Execute() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}
