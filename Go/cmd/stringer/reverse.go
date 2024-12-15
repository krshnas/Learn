package stringer

import (
	"github.com/krishna/go/learn/cli/pkg/stringer"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// reverseCmd reverses the provided string and logs the result.
var reverseCmd = &cobra.Command{
	Use:     "reverse",
	Aliases: []string{"rev"},
	Short:   "Reverses a string",
	Args:    cobra.ExactArgs(1),
	Run:     reverseRun,
}

// reverseRun is the handler function for the "reverse" command.
func reverseRun(cmd *cobra.Command, args []string) {
	// Reverse the input string and log the result
	result := stringer.Reverse(args[0])
	logger.Info("Reversed string", zap.String("original", args[0]), zap.String("reversed", result))
}

// init initializes the reverse command and adds it to the root command.
func init() {
	rootCmd.AddCommand(reverseCmd)
}
