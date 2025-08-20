package uniquenames

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

// Style defines the casing for the generated name.
type Style int

const (
	PascalCase Style = iota
	CamelCase
	SnakeCase
	KebabCase
	LowerCase
	UpperCase
)

// Config holds the configuration for the name generation.
type Config struct {
	Dictionaries [][]string
	Separator    string
	Style        Style
	Seed         int64
}

// Generate creates a unique name based on the provided configuration.
func Generate(config Config) (string, error) {
	if len(config.Dictionaries) == 0 {
		return "", fmt.Errorf("at least one dictionary must be provided")
	}

	words := make([]string, len(config.Dictionaries))
	for i, dict := range config.Dictionaries {
		if len(dict) == 0 {
			return "", fmt.Errorf("dictionary at index %d is empty", i)
		}
		idxBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(dict))))
		if err != nil {
			return "", fmt.Errorf("could not generate random number: %w", err)
		}
		words[i] = dict[idxBig.Int64()]
	}

	result := make([]string, len(words))
	for i, word := range words {
		switch config.Style {
		case PascalCase, CamelCase:
			result[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		case SnakeCase, KebabCase, LowerCase:
			result[i] = strings.ToLower(word)
		case UpperCase:
			result[i] = strings.ToUpper(word)
		}
	}

	finalName := strings.Join(result, config.Separator)

	if config.Style == CamelCase {
		finalName = strings.ToLower(result[0]) + finalName[len(result[0]):]
	}

	return finalName, nil
}
