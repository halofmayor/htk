package main

import (
	"fmt"
	"os"

	"github.com/halofmayor/htk/internal"
)

func main() {
	//Needs a better input validation

	cmd := os.Args[1]
	query := os.Args[2]

	switch cmd {
	case "whatport", "wp":
		fmt.Println(internal.WhatPort(query))
	default:
		fmt.Printf("Command '%s' not recognized\n", cmd)
	}
}
