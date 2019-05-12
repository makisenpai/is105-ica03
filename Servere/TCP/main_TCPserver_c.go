package main

import (
	"encoding/json"
	"log"
	"net"
)

type User struct {
	Name string
	Mail string
}

func handler(c net.Conn, b []byte) {
	c.Write(b)
	c.Close()
}

func main() {

	u := User{"Martin Stenberg", "Martin.v.stenberg@gmail.com"}
	b, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	for {

		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c, b)
	}
}
