package cmd

import (
    "github.com/spf13/cobra"
    "github.com/JorkDev/my-go-cli/config"
    "github.com/JorkDev/my-go-cli/internal"
    "go.uber.org/zap"
    "fmt"
)

var RootCmd = &cobra.Command{
    Use:   "mycli",
    Short: "MyCLI is a simple CLI",
    Run: func(cmd *cobra.Command, args []string) {
        internal.Logger.Info("Root command executed")
        fmt.Println("Welcome to MyCLI")
    },
}

func Execute() {
    internal.InitLogger()
    config.InitConfig()
    if err := RootCmd.Execute(); err != nil {
        internal.Logger.Error("Command execution failed", zap.Error(err))
        fmt.Println("Error:", err)
    }
}
