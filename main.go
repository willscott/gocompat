package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/willscott/gocompat/internal/modfile"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: gocompat <go.mod> <other go.mod>\n")
		os.Exit(2)
	}

	first := args[0]
	second := args[1]
	first_mod := loadMod(first)
	second_mod := loadMod(second)

	for _, r := range first_mod.Require {
		checkMismatch(r, second_mod.Require)
	}
}

func loadMod(mod string) *modfile.File {
	dat, err := ioutil.ReadFile(mod)
	if err != nil {
		panic(err)
	}

	parsed, err := modfile.ParseLax(mod, dat, nil)
	if err != nil {
		panic(err)
	}
	return parsed
}

func checkMismatch(candidate *modfile.Require, group []*modfile.Require) {
	for _, other := range group {
		if candidate.Mod.Path == other.Mod.Path {
			if candidate.Mod.Version != other.Mod.Version {
				fmt.Fprintf(os.Stderr, "Version mismatch in %s: %s vs %s\n", candidate.Mod.Path, candidate.Mod.Version, other.Mod.Version)
				os.Exit(1)
			}
		}
	}
}
