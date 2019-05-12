package main

import (
	"fmt"
	"net"
)

func handler(c net.Conn) {
	c.Write([]byte("ok"))
	c.Close()
}

func main() {
	fmt.Println("Starter Server")
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	for {

		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c)
	}
}
