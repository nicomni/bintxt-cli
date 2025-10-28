package bintxt

import "testing"

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
