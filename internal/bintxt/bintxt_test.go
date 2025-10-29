package bintxt

import (
	"regexp"
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		data string
		want string
	}{
		{
			name: "empty case",
			data: "",
			want: "",
		},
		{
			name: "basic case",
			data: "a b\nc\td",
			want: "01100001 00100000 01100010 00001010 01100011 00001001 01100100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Encode(tt.data)
			if got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		want    string
		wantErr bool
		err     string
	}{
		{
			name: "empty string",
			data: "",
			want: "",
		},
		{
			name: "basic case",
			data: "01100001 00100000 01100010 00001010 01100011 00001001 01100100",
			want: "a b\nc\td",
		},
		{
			name:    "bit length lower than 8",
			data:    "1100001",
			wantErr: true,
			err:     "segment 1: parsing: \"1100001\": syntax: binary value must be exactly 8 bits, was 7 bits",
		},
		{
			name:    "bit length larger than 8",
			data:    "000000001100001",
			wantErr: true,
			err:     "segment 1: parsing: \"000000001100001\": syntax: binary value must be exactly 8 bits, was 15 bits",
		},
		{
			name:    "binary value out of 8-bit range",
			data:    "01100001 100000000",
			wantErr: true,
			err:     "segment 2: parsing: \"100000000\": .*",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Decode(tt.data)

			if gotErr != nil {
				checkError(t, gotErr, tt.wantErr, tt.err)
			}

			if gotErr == nil && tt.wantErr {
				t.Fatalf("Decode() expected error but got none")
			}

			if got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func checkError(t *testing.T, gotErr error, wantErr bool, errPattern string) {
	if gotErr == nil {
		return
	}
	if !wantErr {
		t.Errorf("Decode() unexpected error: %v", gotErr)
	}
	if wantErr && errPattern != "" {
		re := regexp.MustCompile(errPattern)
		if !re.MatchString(gotErr.Error()) {
			t.Errorf("Decode() error = '%v', want pattern '%v'", gotErr.Error(), errPattern)
		}
	}
}
