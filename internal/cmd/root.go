package cmd

import (
	"github.com/nicomni/bintxt-cli/internal/iostreams"
	"github.com/spf13/cobra"
)

func NewCmdRoot(ios *iostreams.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use: "bintxt",
	}

	cmd.AddCommand(NewCmdEncode(ios))
	cmd.AddCommand(NewCmdDecode(ios))
	return cmd
}
