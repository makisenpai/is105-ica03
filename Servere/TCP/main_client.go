package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

type User struct {
	Name string
	Mail string
}

func main() {

	if len(os.Args) == 1 {
		log.Fatal("No ip specified")
	}

	conn, err := net.Dial("tcp", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	b, err := bufio.NewReader(conn).ReadBytes(byte('\n'))
	var u User
	err = json.Unmarshal(b, &u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u.Name)
	fmt.Println(u.Mail)
}
