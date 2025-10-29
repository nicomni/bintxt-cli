package cmd

import (
	"testing"
)

func TestNewCmdRoot(t *testing.T) {
	ios, _, _, _ := setupIO(t, nil)
	cmd := NewCmdRoot(ios)

	t.Run("root command has correct name", func(t *testing.T) {
		wantName := "bintxt"
		if cmd.Name() != wantName {
			t.Fatalf("root.Name() = %q, want: %q", cmd.Name(), wantName)
		}
	})
}

func TestRootHasSubCommands(t *testing.T) {
	ios, _, _, _ := setupIO(t, nil)
	cmd := NewCmdRoot(ios)

	tests := []struct {
		name string
	}{
		{
			name: "encode",
		},
		{
			name: "decode",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, c := range cmd.Commands() {
				if c.Name() == tt.name {
					return
				}
			}
			t.Errorf("subcommand %q not found", tt.name)
		})
	}
}
