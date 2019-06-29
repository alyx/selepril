/*
A very simple TCP server written in Go.
This is a toy project that I used to learn the fundamentals of writing
Go code and doing some really basic network stuff.
Maybe it will be fun for you to read. It's not meant to be
particularly idiomatic, or well-written for that matter.
*/
package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var addr = "localhost"
var port = "9090"

func startServer() {

	fmt.Println("Starting server...")

	src := "localhost:9090"
	listener, _ := net.Listen("tcp", src)
	fmt.Printf("Listening on %s.\n", src)

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Some connection error: %s\n", err)
		}

		go serverHandleConnection(conn)
	}
}

func serverHandleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)

	scanner := bufio.NewScanner(conn)

	for {
		ok := scanner.Scan()

		if !ok {
			break
		}

		serverHandleMessage(scanner.Text(), conn)
	}

	fmt.Println("Client at " + remoteAddr + " disconnected.")
}

func serverHandleMessage(message string, conn net.Conn) {
	msg := strings.TrimSpace(message)
	if len(msg) > 0 {
		if msg != "YDB>" {
			IRC.Privmsg("#alyx", msg)
		}
	}
}
