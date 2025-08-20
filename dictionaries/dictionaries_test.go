package dictionaries

import (
	"strings"
	"testing"
)

// helper function: check for duplicates
func hasDuplicates(list []string) bool {
	seen := make(map[string]bool)
	for _, v := range list {
		if seen[v] {
			return true
		}
		seen[v] = true
	}
	return false
}

// Test adjectives dictionary
func TestAdjectives(t *testing.T) {
	if len(Adjectives) == 0 {
		t.Errorf("Adjectives dictionary is empty")
	}
	if hasDuplicates(Adjectives) {
		t.Errorf("Adjectives dictionary has duplicates")
	}
}

// Test animals dictionary
func TestAnimals(t *testing.T) {
	if len(Animals) == 0 {
		t.Errorf("Animals dictionary is empty")
	}
	if hasDuplicates(Animals) {
		t.Errorf("Animals dictionary has duplicates")
	}
}

// Test cities dictionary
func TestCities(t *testing.T) {
	if len(Cities) == 0 {
		t.Errorf("Cities dictionary is empty")
	}
	if hasDuplicates(Cities) {
		t.Errorf("Cities dictionary has duplicates")
	}
}

// Test colors dictionary
func TestColors(t *testing.T) {
	if len(Colors) == 0 {
		t.Errorf("Colors dictionary is empty")
	}
	if hasDuplicates(Colors) {
		t.Errorf("Colors dictionary has duplicates")
	}
}

// Test countries dictionary
func TestCountries(t *testing.T) {
	if len(Countries) == 0 {
		t.Errorf("Countries dictionary is empty")
	}
	if hasDuplicates(Countries) {
		t.Errorf("Countries dictionary has duplicates")
	}
}

// Test domains dictionary
func TestDomains(t *testing.T) {
	if len(Domains) == 0 {
		t.Errorf("Domains dictionary is empty")
	}
	if hasDuplicates(Domains) {
		t.Errorf("Domains dictionary has duplicates")
	}
}

// Test languages dictionary
func TestLanguages(t *testing.T) {
	if len(Languages) == 0 {
		t.Errorf("Languages dictionary is empty")
	}
	if hasDuplicates(Languages) {
		t.Errorf("Languages dictionary has duplicates")
	}
}

// Test names dictionary
func TestNames(t *testing.T) {
	if len(Names) == 0 {
		t.Errorf("Names dictionary is empty")
	}
	if hasDuplicates(Names) {
		t.Errorf("Names dictionary has duplicates")
	}
}

// Test noun dictionary
func TestNoun(t *testing.T) {
	if len(Noun) == 0 {
		t.Errorf("Noun dictionary is empty")
	}
	if hasDuplicates(Noun) {
		t.Errorf("Noun dictionary has duplicates")
	}
}

// Test programming languages dictionary
func TestProgrammingLanguages(t *testing.T) {
	if len(ProgrammingLanguages) == 0 {
		t.Errorf("ProgrammingLanguages dictionary is empty")
	}
	if hasDuplicates(ProgrammingLanguages) {
		t.Errorf("ProgrammingLanguages dictionary has duplicates")
	}

	// sanity check: entries should not contain spaces at ends
	for _, lang := range ProgrammingLanguages {
		if strings.TrimSpace(lang) != lang {
			t.Errorf("ProgrammingLanguages contains entry with leading/trailing spaces: %q", lang)
		}
	}
}

// Test starwars dictionary
func TestStarWars(t *testing.T) {
	if len(StarWars) == 0 {
		t.Errorf("StarWars dictionary is empty")
	}
	if hasDuplicates(StarWars) {
		t.Errorf("StarWars dictionary has duplicates")
	}
}
