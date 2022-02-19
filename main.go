package main

import (
	"fmt"
	"os"
	"path/filepath"

	"go-wordle/ansi"
	"go-wordle/poorman"
)

func main() {
	if len(os.Args) < 2 {
		if ansi.Main() != 0 {
			fmt.Println("Not ANSI terminal, falling back to poor man option.")
			poorman.Main()
		}
	} else if len(os.Args) > 2 {
		fmt.Printf("usage: %s [OPTION]\n", filepath.Base(os.Args[0]))
		fmt.Println("      OPTON: -p, --poorman")
		fmt.Println("             poor man terminal version (no ANSI)")
	} else {
		if (os.Args[1] == "-p") || (os.Args[1] == "--poorman") {
			poorman.Main()
		} else {
			fmt.Printf("Invalid switch '%s' in command line.\n", os.Args[1])
		}
	}
	os.Exit(0)
}
