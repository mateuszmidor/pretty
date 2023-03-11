package main

import (
	"os"

	"github.com/mateuszmidor/pretty/v2"
)

func main() {
	const INDENT = 4
	pretty.Print(os.Stdout, INDENT)
}
