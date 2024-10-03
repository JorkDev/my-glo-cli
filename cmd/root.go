package cmd

import (
    "github.com/spf13/cobra"
    "github.com/JorkDev/my-go-cli/internal"
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
    internal.InitLogger()
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}
