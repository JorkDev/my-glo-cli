package cmd

import (
	"errors"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestEchoCommandValidInput(t *testing.T) {
	cmd := &cobra.Command{
		Use:   "echo [text]",
		Short: "Echo the input text",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	args := []string{"Hello"}
	cmd.SetArgs(args)

	err := cmd.Execute()
	assert.NoError(t, err, "Expected no error for valid input")
}

func TestEchoCommandInvalidInput(t *testing.T) {
	cmd := &cobra.Command{
		Use:   "echo [text]",
		Short: "Echo the input text",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args[0]) > 100 {
				return errors.New("input text is too long")
			}
			return nil
		},
	}

	longText := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
	args := []string{longText}
	cmd.SetArgs(args)

	err := cmd.Execute()
	assert.Error(t, err, "Expected an error for input longer than 100 characters")
}
