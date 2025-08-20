package uniquenames

import (
	"strings"
	"testing"
)

func TestGenerateErrors(t *testing.T) {
	_, err := Generate(Config{})
	if err == nil {
		t.Errorf("expected error when no dictionaries provided")
	}

	_, err = Generate(Config{
		Dictionaries: [][]string{{}, {"test"}},
	})
	if err == nil {
		t.Errorf("expected error when a dictionary is empty")
	}
}

func TestGenerateWithDifferentStyles(t *testing.T) {
	dictionaries := [][]string{
		{"alpha"},
		{"beta"},
	}

	tests := []struct {
		name     string
		config   Config
		expected func(string) bool // validation function
	}{
		{
			name: "PascalCase",
			config: Config{
				Dictionaries: dictionaries,
				Separator:    "",
				Style:        PascalCase,
			},
			expected: func(out string) bool {
				// Words should start with uppercase
				return strings.HasPrefix(out, "Alpha") || strings.HasPrefix(out, "Beta")
			},
		},
		{
			name: "CamelCase",
			config: Config{
				Dictionaries: dictionaries,
				Separator:    "",
				Style:        CamelCase,
			},
			expected: func(out string) bool {
				// First word lowercase, next word capitalized
				return strings.HasPrefix(out, "alpha") || strings.HasPrefix(out, "beta")
			},
		},
		{
			name: "SnakeCase",
			config: Config{
				Dictionaries: dictionaries,
				Separator:    "_",
				Style:        SnakeCase,
			},
			expected: func(out string) bool {
				return strings.Contains(out, "_") && strings.ToLower(out) == out
			},
		},
		{
			name: "KebabCase",
			config: Config{
				Dictionaries: dictionaries,
				Separator:    "-",
				Style:        KebabCase,
			},
			expected: func(out string) bool {
				return strings.Contains(out, "-") && strings.ToLower(out) == out
			},
		},
		{
			name: "UpperCase",
			config: Config{
				Dictionaries: dictionaries,
				Separator:    "_",
				Style:        UpperCase,
			},
			expected: func(out string) bool {
				return strings.ToUpper(out) == out
			},
		},
		{
			name: "LowerCase",
			config: Config{
				Dictionaries: dictionaries,
				Separator:    ".",
				Style:        LowerCase,
			},
			expected: func(out string) bool {
				return strings.ToLower(out) == out
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Generate(tt.config)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tt.expected(got) {
				t.Errorf("unexpected output for %s: got %s", tt.name, got)
			}
		})
	}
}
