// wacli - WhatsApp CLI tool for sending messages via the WhatsApp Business API
// Fork of steipete/wacli with additional features and improvements
package main

import (
	"fmt"
	"os"

	"github.com/yourusername/wacli/cmd"
)

var (
	// Version is set at build time via ldflags
	Version = "dev"
	// Commit is set at build time via ldflags
	Commit = "none"
	// Date is set at build time via ldflags
	Date = "unknown"
)

func main() {
	cmd.SetVersionInfo(Version, Commit, Date)

	if err := cmd.Execute(); err != nil {
		// Use exit code 2 for usage errors, 1 for general errors
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
