package godaddy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var baseURL = "https://api.godaddy.com/v1/domains/"

// Record defines the records to be updated
type Record struct {
	Data         string `json:"data"` // used to store the IP from where the request is made
	TTL          int32  `json:"ttl"`  // The TTL of the record in seconds
	Name         string `json:"name"` // The name of the record (e.g. "www", "@" or "*")
	TypeOfRecord string `json:"type"` // The type of entry for the record e.g. "TXT" or "A" (default is "TXT")
}

// Godaddy defines the parameters to pass to the server
type Godaddy struct {
	DomainName string    // The domain for which to update the DNS records
	APIKey     string    // The API key for your GoDaddy account
	APISecret  string    // The API key secret for your GoDaddy account
	Records    []*Record // An array of objects that defines the records to update.
}

type ReqObject struct {
	httpMethod string		// "GET" or "PUT"
	URL        string		// The API endpoint
	record     *Godaddy
}

func execute(q ReqObject) {
	url := q.URL

	d, err := json.Marshal(q.record.Records)
	if err != nil {
		log.Fatalf("json.Marshal() failed with '%s'\n", err)
		os.Exit(1)
	}

	body := bytes.NewBuffer(d)

	var req *http.Request
	if q.httpMethod == "GET" {
		req, err = http.NewRequest(q.httpMethod, url, nil)
	} else if q.httpMethod == "PUT" {
		req, err = http.NewRequest(q.httpMethod, url, body)
	}

	if err != nil {
		log.Fatalf("The HTTP request construction failed with error %s\n", err)
		os.Exit(1)
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", fmt.Sprintf("sso-key %s:%s", q.record.APIKey, q.record.APISecret))

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("The HTTP request failed with error %s\n", err)
		os.Exit(1)
	} else {
		if response.StatusCode != 200 {
			log.Fatalf("Request was not successful with error %s.\n", response.Status)
			os.Exit(1)
		}
	}
}

func (gd *Godaddy) SetDNS() {
	q := ReqObject{}
	q.URL = fmt.Sprintf(baseURL+"%s/records/%s/%s", gd.DomainName, gd.Records[0].TypeOfRecord, gd.Records[0].Name)
	q.httpMethod = "PUT"
	q.record = gd
	execute(q)
}

func (gd *Godaddy) GetDNS() {

	q := ReqObject{}
	q.URL = fmt.Sprintf(baseURL+"%s/records/%s/%s", gd.DomainName, gd.Records[0].TypeOfRecord, gd.Records[0].Name)
	q.httpMethod = "GET"
	q.record = gd
	execute(q)
}
