package main

import (
	"fmt"
	"net/http"

	jsonmail "github.com/makisenpai/is105-ica03/Servere/HTTP/json"
)

func handler(w http.ResponseWriter, r *http.Request) {
	msg := jsonmail.Encode("Martin Stenberg", "Martin.v.stenberg@gmail.com")
	fmt.Fprintf(w, string(msg))

}

func main() {
	http.HandleFunc("/connected", handler)
	http.ListenAndServe(":8000", nil)
}
