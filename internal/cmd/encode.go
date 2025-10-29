package cmd

import (
	"fmt"
	"io"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/nicomni/bintxt-cli/internal/bintxt"
	"github.com/nicomni/bintxt-cli/internal/iostreams"
	"github.com/spf13/cobra"
)

func NewCmdEncode(ios *iostreams.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "encode",
		Short: "Convert text to binary text",
		Example: heredoc.Doc(`
			bintxt encode hello # Output: '01101000 01100101 01101100 01101100 01101111'
		`),
		RunE: func(_ *cobra.Command, args []string) error {
			return runCmdEncode(args, ios)
		},
	}
	return cmd
}

func runCmdEncode(args []string, ios *iostreams.IOStreams) error {
	var data string

	if len(args) > 0 {
		data = strings.Join(args, " ")
	} else {
		inData, err := io.ReadAll(ios.In)
		if err != nil {
			return fmt.Errorf("could not read from input stream: %w", err)
		}
		if len(inData) == 0 {
			return nil
		}
		data = string(inData)
	}

	_, err := fmt.Fprintln(ios.Out, bintxt.Encode(data))
	if err != nil {
		return fmt.Errorf("could not write to output stream: %w", err)
	}
	return nil
}
