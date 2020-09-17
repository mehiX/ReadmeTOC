package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mehiX/ReadmeTOC/internal"
)

func main() {
	if 2 != len(os.Args) {
		fmt.Printf("Usage: %s path", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	path := os.Args[1]
	generator := internal.NewGenerator(path)

	generator.Generate()

	fmt.Fprintln(os.Stdout, generator.ToC)

}
