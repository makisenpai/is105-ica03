package lineshift

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func SearchForLineshift(fileToSearch string) {

	// Leser fil inn fil til content
	content, err := ioutil.ReadFile(fileToSearch)
	if err != nil {
		log.Fatal(err)
	}

	// Lagrer innholdet i content som string formatert med % X
	contentAsString := fmt.Sprintf("% X", content)

	// Sjekker hvilken form for linjeskift en fil inneholder
	if strings.Contains(contentAsString, "0D 0A") {
		fmt.Println("Filen inneholder \\r\\n og er formatert for Windows.")
	} else if strings.Contains(contentAsString, "0A") && !strings.Contains(contentAsString, "0D") {
		fmt.Println("Filen inneholder \\n og er formatert for Unix.")
	} else {
		fmt.Println("Filen har ikke linjeskift.")
	}
}
