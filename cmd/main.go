package main

import (
	"log"
	"os"

	gd "github.com/yashhere/godaddyDNS/pkg/godaddy"
)

func main() {
	domainName := os.Args[1]
	validationToken := os.Args[2]
	registrar := os.Args[3]

	tld := gd.GetTLD(domainName)
	sd := gd.GetSubdomain(domainName)

	dnsRecordName := "_acme-challenge"
	if sd != "" {
		dnsRecordName = dnsRecordName + "." + sd
	}

	key, secret, ok := gd.GetGodaddyTokens()
	if !ok {
		log.Fatalf("Environment variables are not set properly. Exiting!")
		os.Exit(1)
	}

	if registrar == "goDaddy" {
		godaddy := gd.Godaddy{}

		godaddy.DomainName = tld
		godaddy.APIKey = key
		godaddy.APISecret = secret
		godaddy.Records = append(godaddy.Records, &gd.Record{
			Data:         validationToken,
			TTL:          600,
			Name:         dnsRecordName,
			TypeOfRecord: "TXT",
		})

		godaddy.SetDNS()
	}
}
