package internal

import (
	"fmt"
	"strconv"
	"strings"
)

//The 0 index represents the service in TCP protocol
//The 1 index represents the service in UDP protocol

var ports = map[int][2]string{
	// Well-known ports
	20:   {"FTP Data", ""},         // TCP
	21:   {"FTP Control", ""},      // TCP
	22:   {"SSH", ""},              // TCP
	23:   {"Telnet", ""},           // TCP
	25:   {"SMTP", ""},             // TCP
	53:   {"DNS", "DNS"},           // TCP/UDP
	67:   {"", "DHCP"},             // UDP
	68:   {"", "DHCP"},             // UDP
	69:   {"TFTP", "TFTP"},         // TCP/UDP
	80:   {"HTTP", ""},             // TCP
	110:  {"POP3", ""},             // TCP
	123:  {"", "NTP"},              // UDP
	137:  {"", "NetBIOS Name"},     // UDP
	138:  {"", "NetBIOS Datagram"}, // UDP
	139:  {"NetBIOS Session", ""},  // TCP
	143:  {"IMAP", ""},             // TCP
	161:  {"", "SNMP"},             // UDP
	162:  {"", "SNMP Trap"},        // UDP
	443:  {"HTTPS", ""},            // TCP
	445:  {"Microsoft-DS", ""},     // TCP
	514:  {"Syslog", "Syslog"},     // TCP/UDP
	636:  {"LDAPS", ""},            // TCP
	993:  {"IMAPS", ""},            // TCP
	995:  {"POP3S", ""},            // TCP
	3306: {"MySQL", ""},            // TCP
	3389: {"RDP", ""},              // TCP
	5060: {"SIP", "SIP"},           // TCP/UDP
	5900: {"VNC", ""},              // TCP
	8080: {"HTTP-Alt", ""},         // TCP

	// Registered ports (1024–49151)
	1025: {"Network Blackjack", ""}, // TCP/UDP
	1080: {"SOCKS Proxy", ""},       // TCP/UDP
	1099: {"RMI Registry", ""},      // TCP
}

var helpText = `whatport is a tool that makes Port <-> Service lookup

whatport supports the following commands:

Usage:
htk whatport <port>        : Returns the service for that port (TCP/UDP)
htk whatport <service>     : Returns the port of that service (TCP/UDP)
htk whatport tcp <port>    : Returns the service for that TCP port
htk whatport udp <port>    : Returns the service for that UDP port
`

func WhatPort(query string) string {
	query = strings.TrimSpace(strings.ToLower(query))
	parts := strings.Fields(query)

	if len(parts) == 0 || parts[0] == "-h" || parts[0] == "--help" {
		return helpText
	}

	index := -1
	if parts[0] == "tcp" {
		index = 0
		parts = parts[1:]
	} else if parts[0] == "udp" {
		index = 1
		parts = parts[1:]
	}

	if len(parts) == 0 {
		return "Please provide a port number or service name."
	}

	target := parts[0]

	if port, err := strconv.Atoi(target); err == nil {
		if svc, ok := ports[port]; ok {
			out := []string{}
			if index == -1 || index == 0 {
				if svc[0] != "" {
					out = append(out, "TCP: "+svc[0])
				}
			}
			if index == -1 || index == 1 {
				if svc[1] != "" {
					out = append(out, "UDP: "+svc[1])
				}
			}
			if len(out) == 0 {
				return "Service not found."
			}
			return strings.Join(out, ", ")
		}
		return "Port not found."
	}

	// Busca reversa: nome do serviço -> porta
	for port, svc := range ports {
		tcpName := strings.ToLower(svc[0])
		udpName := strings.ToLower(svc[1])

		if index == -1 && (tcpName == target || udpName == target) {
			return fmt.Sprintf("Port: %d", port)
		} else if index == 0 && tcpName == target {
			return fmt.Sprintf("Port: %d", port)
		} else if index == 1 && udpName == target {
			return fmt.Sprintf("Port: %d", port)
		}
	}

	return "Service not found."
}
