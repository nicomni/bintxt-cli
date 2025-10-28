package cmd

import (
	"bytes"
	"testing"

	"github.com/google/shlex"
	"github.com/nicomni/bintxt-cli/internal/iostreams"
)

func TestNewCmdEncode(t *testing.T) {
	ios, _, _, _ := setupIO(t, nil)
	cmd := NewCmdEncode(ios)

	t.Run("encode has correct name", func(t *testing.T) {
		expectName := "encode"
		if cmd.Name() != expectName {
			t.Fatalf("encode.Name() = %q, want: %q", cmd.Name(), expectName)
		}
	})
}

func Test_runCmdEncode(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		cliArgs string
		stdin   []byte
		wantOut string
		wantErr bool
	}{
		{
			name:    "single argument",
			cliArgs: "a",
			wantOut: "01100001\n",
		},
		{
			name:    "multiple arguments",
			cliArgs: "a b",
			wantOut: "01100001 00100000 01100010\n",
		},
		{
			name:    "Hello, World! (unquoted)",
			cliArgs: "Hello, World\\!",
			wantOut: "01001000 01100101 01101100 01101100 01101111 00101100 00100000 01010111 01101111 01110010 01101100 01100100 00100001\n",
		},
		{
			name:    "Hello, World! (quoted)",
			cliArgs: "'Hello, World!'",
			wantOut: "01001000 01100101 01101100 01101100 01101111 00101100 00100000 01010111 01101111 01110010 01101100 01100100 00100001\n",
		},
		{
			name:    "multibyte character",
			cliArgs: "é",
			wantOut: "11000011 10101001\n",
		},
		{
			name:    "no arguments reads from stdin",
			stdin:   []byte("hello"),
			wantOut: "01101000 01100101 01101100 01101100 01101111\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ios, _, out, _ := setupIO(t, tt.stdin)
			args := setupArgs(t, tt.cliArgs)

			gotErr := runCmdEncode(args, ios)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("runCmdEncode() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("runCmdEncode() succeeded unexpectedly")
			}

			gotStdut := out.String()

			if gotStdut != tt.wantOut {
				t.Errorf("runCmdEncode() output = %q, want %q", gotStdut, tt.wantOut)
			}
		})
	}
}

func setupArgs(t *testing.T, cliArgs string) []string {
	t.Helper()
	args, err := shlex.Split(cliArgs)
	if err != nil {
		t.Fatalf("could not parse cliArgs %q: %v", cliArgs, err)
	}
	return args
}

func setupIO(t testing.TB, inContent []byte) (ios *iostreams.IOStreams, in, out, errOut *bytes.Buffer) {
	t.Helper()
	ios, in, out, errOut = iostreams.InMemoryIO()
	in.Write(inContent)
	return
}
