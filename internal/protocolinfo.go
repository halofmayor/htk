package internal

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"
)

//go:embed protocolinfo/protocolJSONs/*/*.JSON
var protocolFiles embed.FS

type Handshake struct {
	Description string   `json:"description"`
	Steps       []string `json:"steps"`
}

type Field struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Message struct {
	Unit      string            `json:"unit"`
	Structure map[string]string `json:"structure"`
}

type Protocol struct {
	Name      string    `json:"name"`
	ShortName string    `json:"shortName"`
	Summary   string    `json:"summary"`
	Layer     string    `json:"layer"`
	Ports     []int     `json:"ports"`
	RFCs      []string  `json:"rfcs"`
	Purpose   string    `json:"purpose"`
	Handshake Handshake `json:"handshake"`
	Fields    []Field   `json:"fields"`
	Message   Message   `json:"message"`
	Usage     []string  `json:"usage"`
	Pitfalls  []string  `json:"pitfalls"`
}

var protocolCategories = map[string]string{
	"TCP":  "datatransport",
	"UDP":  "datatransport",
	"SCTP": "datatransport",
	"DCCP": "datatransport",
	"QUIC": "datatransport",

	"HTTP":   "application",
	"HTTPS":  "application",
	"FTP":    "application",
	"FTPS":   "application",
	"SFTP":   "application",
	"TFTP":   "application",
	"SMTP":   "application",
	"POP3":   "application",
	"IMAP":   "application",
	"DNS":    "application",
	"DNSSEC": "application",
	"NFS":    "application",
	"SMB":    "application",
	"CIFS":   "application",
	"SNMP":   "application",
	"MQTT":   "application",
	"AMQP":   "application",
	"XMPP":   "application",

	"IPSEC":     "securityandcryptography",
	"IKE":       "securityandcryptography",
	"TLS":       "securityandcryptography",
	"SSL":       "securityandcryptography",
	"SSH":       "securityandcryptography",
	"KERBEROS":  "securityandcryptography",
	"RADIUS":    "securityandcryptography",
	"DIAMETER":  "securityandcryptography",
	"OAUTH":     "securityandcryptography",
	"OIDC":      "securityandcryptography",
	"PGP":       "securityandcryptography",
	"SMIME":     "securityandcryptography",
	"SRTP":      "securityandcryptography",
	"WIREGUARD": "securityandcryptography",
	"DTLS":      "securityandcryptography",
}

var helpTextPI = `protocolinfo is a tool that returns information about a certain protocol

protocolinfo supports the following commands:

Usage:
htk protocolinfo <protocol>       : Returns information about the protocol
htk protocolinfo -v <protocol>    : Returns all information about the protocol
htk protocolinfo -o <protocol>    : Returns only the essential information about the protocol
`

func ProtocolInfo(query string) string {
	query = strings.TrimSpace(strings.ToLower(query))
	parts := strings.Fields(query)

	if len(parts) == 0 || parts[0] == "-h" || parts[0] == "--help" {
		return helpTextPI
	}

	showOnlyBasic := false
	showVerbose := false
	protocolName := ""

	for _, part := range parts {
		switch part {
		case "-o":
			showOnlyBasic = true
		case "-v":
			showVerbose = true
		default:
			if part != "pi" && part != "protocolinfo" {
				protocolName = part
			}
		}
	}

	if protocolName == "" {
		return "Error: protocol name not provided.\nUsage: pi [options] <protocol>"
	}

	category, ok := protocolCategories[strings.ToUpper(protocolName)]
	if !ok {
		return fmt.Sprintf("Error: unknown protocol '%s'", protocolName)
	}

	filePath := fmt.Sprintf("protocolJSONs/%s/%s.JSON", category, strings.ToUpper(protocolName))
	data, err := protocolFiles.ReadFile(filePath)
	if err != nil {
		return fmt.Sprintf("Error reading embedded file %s: %v", filePath, err)
	}

	var protocol Protocol
	if err := json.Unmarshal(data, &protocol); err != nil {
		return fmt.Sprintf("Error parsing JSON: %v", err)
	}

	if showOnlyBasic {
		return fmt.Sprintf("Name: %s\nShort Name: %s\nSummary: %s\n", protocol.Name, protocol.ShortName, protocol.Summary)
	} else if showVerbose {
		b, _ := json.MarshalIndent(protocol, "", "  ")
		return string(b)
	} else {
		fieldsSummary := []string{}
		for _, f := range protocol.Fields {
			fieldsSummary = append(fieldsSummary, f.Name)
		}
		return fmt.Sprintf(
			"Name: %s\nShort Name: %s\nSummary: %s\nLayer: %s\nPorts: %v\nRFCs: %v\nPurpose: %s\nUsage: %v\nPitfalls: %v\nHandshake: %s\nFields: %s\nMessage Unit: %s\n",
			protocol.Name, protocol.ShortName, protocol.Summary, protocol.Layer, protocol.Ports, protocol.RFCs, protocol.Purpose,
			protocol.Usage, protocol.Pitfalls, protocol.Handshake.Description, strings.Join(fieldsSummary, ", "), protocol.Message.Unit)
	}
}
