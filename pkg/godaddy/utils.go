package godaddy

import (
	"os"
	"strings"
)

func GetGodaddyTokens() (string, string, bool) {
	key := os.Getenv("GODADDY_KEY")
	secret := os.Getenv("GODADDY_SECRET")
	ok := true

	if key == "" || secret == "" {
		ok = false
	}
	return key, secret, ok
}

func GetTLD(domainName string) string {
	s := strings.Split(domainName, ".")
	if len(s) == 2 {
		// domainName = yashagarwal.in
		// s = ["yashagarwal", "in"]
		return domainName
	}

	// domainName = www.photos.yashagarwal.in
	// s = ["www", "photos", "yashagarwal", "in"]
	return strings.Join(s[len(s)-2:], ".")
}

func GetSubdomain(domainName string) string {
	s := strings.Split(domainName, ".")
	if len(s) == 2 {
		// domainName = yashagarwal.in
		// s = ["yashagarwal", "in"]
		return ""
	}

	// domainName = www.photos.yashagarwal.in
	// s = ["www", "photos", "yashagarwal", "in"]
	return strings.Join(s[:len(s)-2], ".")
}
