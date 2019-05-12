package main

import (
	"os"

	"github.com/makisenpai/is105-ica03/fileinfo"
)

func main() {

	if len(os.Args) < 3 {

		return

	}

	if os.Args[1] != "-f" {

		return

	}

	fileinfo.ReadFileinfo(os.Args[2])

}
