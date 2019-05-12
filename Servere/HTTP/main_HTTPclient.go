package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	re, err := http.Get("http://localhost:8000/connected")
	if err != nil {
		log.Fatal(err)
	}
	defer re.Body.Close()

	body, err := ioutil.ReadAll(re.Body)
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stdout.Write(body)
	if err != nil {
		log.Fatal(err)
	}
}
