package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	content1, err := ioutil.ReadFile("files/text1.txt")
	if err != nil {
		log.Fatal(err)
	}

	content2, err2 := ioutil.ReadFile("files/text2.txt")
	if err2 != nil {
		log.Fatal(err)
	}

	fmt.Printf("Fil 1 inneholder: %s", content1)
	fmt.Printf("Fil 2 inneholder: %s", content2)

	fmt.Println(bytes.Equal(content1, content2))

	fmt.Printf("Fil 1 lengde: %v\nFil 2 lengde: %v\n", len(content1), len(content2))

	for a := range content2 {
		fmt.Printf("%q , %q\n", content1[a], content2[a])
	}

}
