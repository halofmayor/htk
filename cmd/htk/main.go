package main

import (
	"fmt"
	"os"
	"strings"

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

	args := os.Args[1:] // ignora "htk"
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		fmt.Println("HTK - Halof ToolKit\n\nUse 'htk whatport -h' for help on the whatport module.")
		return
	}

	switch args[0] {
	case "whatport", "wp":
		query := ""
		if len(args) > 1 {
			query = args[1:]
			queryStr := ""
			for i := 1; i < len(args); i++ {
				queryStr += args[i] + " "
			}
			query = queryStr
		}
		fmt.Println(internal.WhatPort(query))
	default:
		fmt.Println("Unknown command. Use 'htk -h'")
	}
}
