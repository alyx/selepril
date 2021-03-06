package main

import (
	"crypto/tls"
	"fmt"
	"strings"
	"sync"
	"time"

	irc "github.com/thoj/go-ircevent"
)

var wg sync.WaitGroup

// IRC IRC Conn
var IRC *irc.Connection

func handlePrivmsg(e *irc.Event) {
	rawmsg := e.Message()
	if strings.HasPrefix(rawmsg, "^add") {
		msg := strings.Replace(rawmsg, "^add", "", 1)
		msg = strings.Replace(msg, "\"", "'", 5000)
		msg = strings.TrimSpace(msg)
		CConn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		CConn.Write([]byte("DO ADD^Quotes(" + msg + ")\n"))
		IRC.Privmsg(e.Source, "Quote added")
	}

	if len(e.Arguments) > 1 {
		if e.Arguments[0] == "^get" {
			CConn.Write([]byte("DO READ^Quotes(" + strings.Replace(e.Arguments[1], "\"", "'", 5000) + ")"))
		} else if e.Arguments[0] == "^rand" {
			CConn.Write([]byte("DO RAND^Quotes()"))
		}
	}
}

func main() {
	ircnick1 := "Mallory"
	irccon := irc.IRC(ircnick1, "mallory")
	irccon.VerboseCallbackHandler = true
	irccon.Debug = true
	irccon.UseTLS = true
	irccon.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	irccon.AddCallback("001", func(e *irc.Event) { irccon.Join("#alyx") })
	irccon.AddCallback("366", func(e *irc.Event) {})
	irccon.AddCallback("PRIVMSG", handlePrivmsg)
	err := irccon.Connect("irc.interlinked.me:6697")
	if err != nil {
		fmt.Printf("Err %s", err)
		return
	}
	IRC = irccon
	wg.Add(3)
	go startServer()
	fmt.Println("Proceed when out server online")
	time.Sleep(15)
	go startClient()
	go irccon.Loop()
	wg.Wait()
}
