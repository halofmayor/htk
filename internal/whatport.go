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

var services = func() map[[2]string][]int {
	m := make(map[[2]string][]int)
	for port, svc := range ports {
		m[svc] = append(m[svc], port)
	}
	return m
}()

func WhatPort(query string) string {
	query = strings.ToLower(query)
	index := -1 //-1 Means both tcp and udp

	if strings.Contains(query, "tcp") && !strings.Contains(query, "udp") {
		index = 0
	} else if strings.Contains(query, "udp") && !strings.Contains(query, "tcp") {
		index = 1
	}

	//Removes the tcp and upd from the query if it exists
	query = strings.ReplaceAll(query, "tcp", "")
	query = strings.ReplaceAll(query, "udp", "")
	query = strings.TrimSpace(query)

	//If port
	if port, err := strconv.Atoi(query); err == nil {
		if svc, ok := ports[port]; ok {
			if index == -1 {
				return fmt.Sprintf("TCP: %s, UDP: %s", svc[0], svc[1])
			} else if index == 0 {
				return fmt.Sprintf("TCP: %s", svc[0])
			} else {
				return fmt.Sprintf("UDP: %s", svc[1])
			}
		} else {
			return "Port not found."
		}
	} else {
		for port, svc := range ports {
			if (index == -1 && strings.ToLower(svc[0]) == query || strings.ToLower(svc[1]) == query) ||
				(index == 0 && strings.ToLower(svc[0]) == query) ||
				(index == 1 && strings.ToLower(svc[1]) == query) {
				return fmt.Sprintf("Porta: %d", port)
			}
		}
		return "Service not found."
	}
}
