package main

import (
	"os"

	"github.com/makisenpai/is105-ica03/lineshift"
)

func main() {
	filename := os.Args[1]
	lineshift.SearchForLineshift(filename)
}
