package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hmarr/codeowners"
)

func main() {
	file, err := os.Open("/src/CODEOWNERS")
	if err != nil {
		log.Fatal(err)
	}

	ruleset, err := codeowners.ParseFile(file)
	if err != nil {
		log.Fatal(err)
	}

	rule, err := ruleset.Match("/src/go.sum")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Owners: %v\n", rule.Owners)
}
