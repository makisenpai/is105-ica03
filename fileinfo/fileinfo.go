package fileinfo

import (
	"fmt"
	"log"
	"os"
)

func ReadFileinfo(filename string) {

	//filenavn
	fi, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	file := fi.Mode()
	fmt.Printf("Info about file %v:\n", fi.Name())
	b := fi.Size()
	KiB := b / 1024
	MiB := KiB / 1024
	GiB := MiB / 1024

	fmt.Printf("Size: %v bytes, %v KiB, %v Mib, %v GiB\n", b, KiB, MiB, GiB)

	if file.IsDir() {
		fmt.Println("Is a directory")
	} else {
		fmt.Println("Is not a directory")
	}

	if file.IsRegular() {
		fmt.Println("Is a regular file")
	} else {
		fmt.Println("Is not a regular file")
	}

	fmt.Printf("Has Unix permission bits: %v\n", file.Perm)

	if file&os.ModeAppend != 0 {
		fmt.Println("Is append only")
	} else {
		fmt.Println("Is not append only")
	}

	if file&os.ModeDevice != 0 {
		fmt.Println("Is a device file")
	} else {
		fmt.Println("Is not a device file")
	}

	if file&os.ModeCharDevice != 0 {
		fmt.Println("Is a Unix character device")
	} else {
		fmt.Println("Is not a Unix character device")
	}

	if file&os.ModeSymlink != 0 {
		fmt.Println("Is a symbolic link")
	} else {
		fmt.Println("Is not a symbolic link")
	}

}
