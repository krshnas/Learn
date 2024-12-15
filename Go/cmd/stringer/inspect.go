package stringer

import (
	"github.com/krishna/go/learn/cli/pkg/stringer"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var onlyDigits bool

// inspectCmd inspects the provided string and logs the results.
var inspectCmd = &cobra.Command{
	Use:     "inspect",
	Aliases: []string{"insp"},
	Short:   "Inspects a string",
	Args:    cobra.ExactArgs(1),
	Run:     inspectRun,
}

func inspectRun(cmd *cobra.Command, args []string) {
	// Extract the input string and process the inspection
	inputString := args[0]
	result, kind := stringer.Inspect(inputString, onlyDigits)

	// Determine the plural form for the result
	pluralS := "s"
	if result == 1 {
		pluralS = ""
	}

	// Log the inspection result
	logger.Info("Stringer Inspect", zap.String("item", inputString), zap.Int("result", result), zap.String("kind", kind), zap.String("plural", pluralS))
}

// init initializes the inspect command, adding flags and configuring it.
func init() {
	// Add the "digits" flag to the inspect command
	inspectCmd.Flags().BoolVarP(&onlyDigits, "digits", "d", false, "Count only digits")

	// Register the inspect command with the root command
	rootCmd.AddCommand(inspectCmd)
}
