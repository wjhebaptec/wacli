package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version information set by goreleaser at build time via ldflags.
var (
	Version   = "dev"
	Commit    = "none"
	BuildDate = "unknown"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Long:  `Print the version, commit hash, and build date of wacli.`,
	Run:   runVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

// runVersion prints version details to stdout.
func runVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("wacli version %s\n", Version)
	fmt.Printf("  commit:     %s\n", Commit)
	fmt.Printf("  built:      %s\n", BuildDate)
}
