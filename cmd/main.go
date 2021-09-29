package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dvoros/pocketmine-rcon/pkg/connection"
	"github.com/sirupsen/logrus"
)

func main() {
	if os.Getenv("RCON_DEBUG") == "true" {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if len(os.Args) < 4 {
		logrus.Fatal("Usage: ./rcon address password command(s)")
	}
	addr := os.Args[1]
	pass := os.Args[2]
	command := strings.Join(os.Args[3:], " ")
	logrus.Debugf("Command to run: %s", command)

	conn, err := connection.NewConnection(addr, pass)
	if err != nil {
		fmt.Println(err)
		return
	}
	logrus.Debugf("Successfully logged in at %s", addr)

	r, err := conn.SendCommand(command)
	if err != nil {
		logrus.Fatalf("Error: %s", err)
	}

	fmt.Print(r)
}
