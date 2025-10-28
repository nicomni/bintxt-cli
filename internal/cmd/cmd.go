// Package cmd contains the command-line interface implementation for the
// application bintxt.
package cmd

import (
	"fmt"

	"github.com/nicomni/bintxt-cli/internal/iostreams"
)

type ExitCode int

const (
	ExitSuccess ExitCode = 0
	ExitError   ExitCode = 1
)

func Main() ExitCode {
	ios := iostreams.SystemIO()
	rootCmd := NewCmdRoot(ios)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("Error:", err)
		return ExitError
	}
	return ExitSuccess
}
