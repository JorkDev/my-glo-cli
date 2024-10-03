package cmd

import (
    "errors"
    "fmt"
    "github.com/spf13/cobra"
    "go.uber.org/zap"
    "github.com/JorkDev/my-go-cli/internal"
)

var echoCmd = &cobra.Command{
    Use:   "echo [text]",
    Short: "Echo the input text",
    Args:  cobra.MinimumNArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        if len(args[0]) > 100 {
            internal.Logger.Error("Text length exceeds 100 characters", zap.String("input", args[0]))
            return errors.New("input text is too long, must be under 100 characters")
        }

        fmt.Println(args[0])
        return nil
    },
}

func init() {
    RootCmd.AddCommand(echoCmd)
}
