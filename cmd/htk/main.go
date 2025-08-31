package main

import (
	"fmt"
	"os"

	"github.com/halofmayor/htk/internal"
)

func Help() {
	helpText := `whatport supports the following available commands:

    whatport <port>            : Returns the service in that port. (TCP and UDP)
    whatport <service>         : Returns the port that the service is running. (TCP and UDP)
    whatport tcp <port>        : Returns the service in that port. (TCP only)
    whatport udp <port>        : Returns the service in that port. (UDP only)
    whatport tcp <service>     : Returns the port that the service is running. (TCP only)
    whatport udp <service>     : Returns the port that the service is running. (UDP only)

    Options:
    -h, --help                 : Show this help message
    `
	fmt.Println(helpText)
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Missing command. Use: htk -h for help")
		os.Exit(1)
	}

	// Help geral
	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println(`HTK - Toolset
					Usage:
					htk <command> [options]

					Available commands:
					whatport : Check ports and services.
					`)
		return
	}

	cmd := os.Args[1]
	query := os.Args[2]

	switch cmd {
	case "-h", "--help":
		Help()
		return
	case "whatport", "wp":
		fmt.Println(internal.WhatPort(query))
		return
	default:
		fmt.Printf("Command '%s' not recognized.\n", cmd)
	}
}
