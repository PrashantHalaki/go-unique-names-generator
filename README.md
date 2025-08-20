# go-unique-names-generator

A flexible and customizable Go package to generate unique, human-readable names from multiple dictionaries.
Useful for creating usernames, project names, temporary IDs, or any situation where you need **unique but meaningful strings**.

---

## ✨ Features

- Combine words from multiple dictionaries (adjectives, animals, colors, cities, etc.)
- Multiple output styles: **PascalCase, CamelCase, snake_case, kebab-case, UPPERCASE, lowercase**
- Configurable **separators** (`-`, `_`, `.` or anything you choose)
- Randomized output powered by crypto-safe randomness
- Easy to extend with your own dictionaries

---

## 📦 Installation

```bash
go get github.com/PrashantHalaki/go-unique-names-generator
```

---

## 🚀 Usage

```go
package main

import (
	"fmt"

	"github.com/PrashantHalaki/go-unique-names-generator"
	"github.com/PrashantHalaki/go-unique-names-generator/dictionaries"
)

func main() {
	// Example configuration
	config := uniquenames.Config{
		Dictionaries: [][]string{
			dictionaries.Adjectives,
			dictionaries.Animals,
		},
		Separator: "-",
		Style:     uniquenames.PascalCase,
	}

	// Generate names
	for i := 0; i < 5; i++ {
		name, err := uniquenames.Generate(config)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println("Generated name:", name)
	}
}
```

---

## 🔥 Example Output

```
Brave-Tiger
Happy-Lion
Swift-Elephant
Clever-Cat
Mighty-Wolf
```

---

## 📚 Available Dictionaries

Located in the `dictionaries/` folder:

- `adjectives.go`
- `animals.go`
- `cities.go`
- `colors.go`
- `countries.go`
- `domains.go`
- `languages.go`
- `names.go`
- `noun.go`
- `programminglanguages.go`
- `starwars.go`

You can also create your own dictionary `.go` file and include it in the `Config`.

---

## 🧪 Running Tests

```bash
go test ./...
```

---

## 🛠️ Custom Dictionary Example

```go
myWords := []string{"alpha", "beta", "gamma"}

config := uniquenames.Config{
	Dictionaries: [][]string{myWords},
	Separator:    "_",
	Style:        uniquenames.UpperCase,
}

name, _ := uniquenames.Generate(config)
fmt.Println(name) // ALPHA or BETA or GAMMA
```

---

## 🤝 Contributing

Contributions are welcome!
You can:

- Add more dictionaries
- Suggest new output styles
- Improve test coverage

Just fork the repo, make changes, and submit a PR.

---

## 📄 License

MIT License © 2025 [Prashant Halaki](https://github.com/PrashantHalaki)
