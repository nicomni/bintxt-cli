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

func NewCmdDecode(ios *iostreams.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decode",
		Short: "Convert binary text to readable text",
		Example: heredoc.Doc(`
			bintxt decode 01100001 # Output: 'a'
		`),
		RunE: func(_ *cobra.Command, args []string) error {
			return runCmdDecode(args, ios)
		},
	}

	return cmd
}

func runCmdDecode(args []string, ios *iostreams.IOStreams) error {
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
	result, err := bintxt.Decode(data)
	if err != nil {
		return fmt.Errorf("could not decode input: %w", err)
	}

	_, err = fmt.Fprintln(ios.Out, result)
	if err != nil {
		return fmt.Errorf("could not write to output stream: %w", err)
	}
	return nil
}
