package main

import (
	"os"

	"github.com/nicomni/bintxt-cli/internal/cmd"
)

func main() {
	code := cmd.Main()
	os.Exit(int(code))
}
