package main

import (
	"github.com/arnekd/is105-ica03/lineshift"
	"os"
)

func main() {
	filename := os.Args[1]
	lineshift.SearchForLineshift(filename)
}
