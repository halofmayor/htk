package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/halofmayor/htk/internal"
)

func Help() {
	helpText := `Halof ToolKit supports the following commands:

whatport <port>            : Returns the service in that port. (TCP and UDP)
whatport <service>         : Returns the port that the service is running. (TCP and UDP)
whatport tcp <port>        : Returns the service in that port. (TCP only)
whatport udp <port>        : Returns the service in that port. (UDP only)
whatport tcp <service>     : Returns the port that the service is running. (TCP only)
whatport udp <service>     : Returns the port that the service is running. (UDP only)

protocolinfo <protocol> : Shows information about the protocol.
protocolinfo -v <protocol> : Shows all information about the protocol.
protocolinfo -o <protocol> : Shows only the important information about the protocol.

Options:
-h, --help                 : Show this help message
`
	fmt.Println(helpText)
}

func main() {

	args := os.Args[1:]
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		Help()
		return
	}

	switch args[0] {
	case "whatport", "wp":
		query := ""
		if len(args) > 1 {
			query = strings.Join(args[1:], " ")
		}
		fmt.Println(internal.WhatPort(query))
	case "protocolinfo", "pi":
		query := ""
		if len(args) > 1 {
			query = strings.Join(args[1:], " ")
		}
		fmt.Println(internal.ProtocolInfo(query))
	default:
		fmt.Println("Unknown command. Use 'htk -h'")
	}

}
