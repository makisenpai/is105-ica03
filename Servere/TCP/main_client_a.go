package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("Noe funker ikke :(")

	}
	fmt.Fprintf(conn, "GET /HTTP/1.0\n\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)

}
