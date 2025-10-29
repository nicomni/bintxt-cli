package cmd

import "testing"

func TestNewCmdDecode(t *testing.T) {
	ios, _, _, _ := setupIO(t, nil)
	cmd := NewCmdDecode(ios)

	t.Run("command has correct name", func(t *testing.T) {
		expectName := "decode"
		if cmd.Name() != expectName {
			t.Fatalf("encode.Name() = %q, want: %q", cmd.Name(), expectName)
		}
	})
}

func Test_runCmdDecode(t *testing.T) {
	tests := []struct {
		name    string
		cliArgs string
		stdin   []byte
		wantOut string
		wantErr bool
	}{
		{
			name:    "basic",
			cliArgs: "01100001",
			wantOut: "a\n",
		},
		{
			name:    "multiple args",
			cliArgs: "01100001 00100000 01100010",
			wantOut: "a b\n",
		},
		{
			name:    "quoted arg",
			cliArgs: "\"01100001 00100000 01100010\"",
			wantOut: "a b\n",
		},
		{
			name:    "no arguments read from stdin",
			stdin:   []byte("01100101 01111000 01100001 01101101 01110000 01101100 01100101"),
			wantOut: "example\n",
		},
		{
			name:    "multi-byte character",
			stdin:   []byte("11000011 10101001"),
			wantOut: "é\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ios, _, out, _ := setupIO(t, tt.stdin)
			args := setupArgs(t, tt.cliArgs)

			gotErr := runCmdDecode(args, ios)

			if gotErr != nil {
				checkError(t, gotErr, tt.wantErr)
				return
			}
			if tt.wantErr {
				t.Fatalf("runCmdDecode() succeeded unexpectedly")
			}

			gotStdut := out.String()

			if gotStdut != tt.wantOut {
				t.Errorf("runCmdDecode() output = %q, want %q", gotStdut, tt.wantOut)
			}
		})
	}
}

func checkError(t testing.TB, gotErr error, wantErr bool) {
	if gotErr == nil {
		return
	}

	if !wantErr {
		t.Errorf("runCmdDecode() failed unexpectedly: %v", gotErr)
	}
}
