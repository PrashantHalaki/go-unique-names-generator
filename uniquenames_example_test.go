package uniquenames

import (
	"fmt"

	"github.com/PrashantHalaki/go-unique-names-generator/dictionaries"
)

// ExampleGenerate demonstrates using the Generate function.
func ExampleGenerate() {
	cfg := Config{
		Dictionaries: [][]string{dictionaries.Adjectives, dictionaries.Cities},
		Separator:    ".",
		Style:        KebabCase,
	}
	name, _ := Generate(cfg)

	// Print Random name
	fmt.Println(name)
}
