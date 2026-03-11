package logic

import "testing"

func TestIsValidURL(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "valid https url",
			input: "https://example.com",
			want:  true,
		},
		{
			name:  "valid http url",
			input: "http://example.com/page",
			want:  true,
		},
		{
			name:  "missing scheme",
			input: "example.com",
			want:  false,
		},
		{
			name:  "invalid text",
			input: "not-a-url",
			want:  false,
		},
		{
			name:  "unsupported scheme",
			input: "ftp://example.com",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidURL(tt.input)
			if got != tt.want {
				t.Fatalf("expected %v, got %v", tt.want, got)
			}
		})
	}
}