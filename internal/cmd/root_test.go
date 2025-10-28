package cmd

import "testing"

func TestNewCmdRoot(t *testing.T) {
	ios, _, _, _ := setupIO(t, nil)
	cmd := NewCmdRoot(ios)

	t.Run("root command has correct name", func(t *testing.T) {
		wantName := "bintxt"
		if cmd.Name() != wantName {
			t.Fatalf("root.Name() = %q, want: %q", cmd.Name(), wantName)
		}
	})

	t.Run("has subcommand: encode", func(t *testing.T) {
		cmds := cmd.Commands()
		for _, c := range cmds {
			if c.Name() == "encode" {
				return
			}
		}
		t.Fatal("encode is not a subcommand of the root")
	})
}
