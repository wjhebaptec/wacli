// Package cmd provides the CLI command structure for wacli.
// wacli is a command-line interface for interacting with WhatsApp Web API.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Version is set at build time via ldflags.
	Version = "dev"
	// Commit is the git commit hash set at build time.
	Commit = "none"
	// Date is the build date set at build time.
	Date = "unknown"
)

// rootCmd is the base command for the wacli application.
var rootCmd = &cobra.Command{
	Use:   "wacli",
	Short: "A command-line interface for WhatsApp Web",
	Long: `wacli is a CLI tool for sending and receiving WhatsApp messages
using the WhatsApp Web multi-device API.

Authenticate once using a QR code, then send messages, files,
and media directly from your terminal or scripts.

Config directory defaults to $HOME/.config/wacli (override with --config or WACLI_CONFIG env var).`,
	SilenceUsage:  true,
	SilenceErrors: true, // handle errors ourselves for cleaner output
}

// versionCmd prints the current version information.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("wacli %s (commit: %s, built: %s)\n", Version, Commit, Date)
	},
}

// Execute runs the root command and exits on error.
func Execute(version, commit, date string) {
	Version = version
	Commit = commit
	Date = date

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Persistent flags available to all subcommands.
	rootCmd.PersistentFlags().StringP(
		"config", "c", "",
		"path to config directory (default: $HOME/.config/wacli, or $WACLI_CONFIG)",
	)
	rootCmd.PersistentFlags().BoolP(
		"verbose", "v", false,
		"enable verbose/debug logging",
	)
	// Default to showing help when no subcommand is provided, rather than
	// silently doing nothing — makes the tool more discoverable for first-time use.
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	}
}
