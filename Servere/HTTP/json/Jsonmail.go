package jsonmail

import (
	"encoding/json"
)

type nameAndEmail struct {
	Email string
	Name  string
}

func getName_Email(name, email string) nameAndEmail {

	response := nameAndEmail{
		Name:  name,
		Email: email}
	return response
}

func Encode(name, email string) []byte {
	input := getName_Email(name, email)
	output, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}
	return output
}
