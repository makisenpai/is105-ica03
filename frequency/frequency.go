package frequence 

import (
	"fmt"
	"io/ioutil"
)

func FileCount(filename string) string {
	file, err := ioutil.ReadFile(filename)
	
	if err != nil{ //vis filen ikke kan åpnes stopper programet.
		panic(err)
	}
	
	NewLCounter := 0 

	for _: b := range file {
		if b == 10 {
			newLCounter++
		}
	}
	return fmt.Sprintf("Number of lines: %v", newLCounter)
}