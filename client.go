package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

// CConn Client Connector
var CConn net.Conn

func startClient() {
	dest := "localhost:9091"
	fmt.Printf("Connecting to %s...\n", dest)

	conn, err := net.Dial("tcp", dest)

	if err != nil {
		if _, t := err.(*net.OpError); t {
			fmt.Println("Some problem connecting.")
		} else {
			fmt.Println("Unknown error: " + err.Error())
		}
		os.Exit(1)
	}
	CConn = conn

	for {
		}
	}
}
