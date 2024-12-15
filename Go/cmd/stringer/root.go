package stringer

import (
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const version = "0.0.1"

var logger *zap.Logger

// rootCmd represents the base command when executed without any subcommands.
var rootCmd = &cobra.Command{
	Use:     "stringer",
	Version: version,
	Short:   "stringer - a simple CLI to transform and inspect strings",
	Long: `stringer is a super fancy CLI (kidding).
   
One can use stringer to modify or inspect strings directly from the terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Placeholder for the command logic
	},
}

// Initialize and execute the root command with proper logging.
func Execute(log *zap.Logger) {
	logger = log
	if err := rootCmd.Execute(); err != nil {
		// Log error with additional context
		logger.Error("Whoops. There was an error while executing your CLI", zap.Error(err), zap.Any("stderr", os.Stderr))
		// Exit with status code 1
		os.Exit(1)
	}
}

// SetupRootCmd initializes the root command with additional flags or functionality if needed.
func SetupRootCmd() {
	// Additional setup for the root command can be done here if needed
}
