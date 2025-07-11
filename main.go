package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	//it uses this nice whois library, interfacing with global whois servers, packages it to serve JSON
	"github.com/likexian/whois"
)

type WhoisResponse struct {
	Target  string      `json:"target"`   //the target IP address we lookup
	Result  interface{} `json:"result"`   //whois information
	ASN     interface{} `json:"asn"`      //ASN number
	ASNName interface{} `json:"asn_name"` //ASN Name
}

// to filter out private IPs, use this functions that recognises private ranges to avoid useless HTTP Requests
func isPrivateIP(ip net.IP) bool {
	privateBlocks := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"169.254.0.0/16",
		"127.0.0.0/8",
		"100.64.0.0/10",
		"192.0.0.0/24",
		"192.0.2.0/24",
		"192.168.0.0/16",
	}
	for _, block := range privateBlocks {
		_, cidr, _ := net.ParseCIDR(block)
		if cidr.Contains(ip) {
			return true
		}
	}
	return false
}

// parses the whois lookup result from github.com/likexian/whois, to provide JSON response
func parseWhois(raw string) map[string]interface{} {
	result := make(map[string]interface{})
	lines := strings.Split(raw, "\n")
	for _, line := range lines {
		if strings.Contains(line, ":") {
			parts := strings.SplitN(line, ":", 2)
			key := strings.TrimSpace(strings.ToLower(parts[0]))
			value := strings.TrimSpace(parts[1])
			if existing, ok := result[key]; ok {
				switch v := existing.(type) {
				case []string:
					result[key] = append(v, value)
				case string:
					result[key] = []string{v, value}
				}
			} else {
				result[key] = value
			}
		}
	}
	return result
}

// ASN lookup using bgpview.io
func getASNInfo(ip string) (interface{}, interface{}) {
	resp, err := http.Get("https://api.bgpview.io/ip/" + ip)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, nil
	}

	if data["status"] != "ok" {
		return nil, nil
	}

	dataField, ok := data["data"].(map[string]interface{})
	if !ok {
		return nil, nil
	}

	prefixes, ok := dataField["prefixes"].([]interface{})
	if !ok || len(prefixes) == 0 {
		return nil, nil
	}

	prefix0, ok := prefixes[0].(map[string]interface{})
	if !ok {
		return nil, nil
	}

	asnData, ok := prefix0["asn"].(map[string]interface{})
	if !ok {
		return nil, nil
	}

	return asnData["asn"], asnData["name"]
}

func lookup(target string) WhoisResponse {
	ip := net.ParseIP(target)
	if ip == nil {
		ips, err := net.LookupIP(target)
		if err != nil || len(ips) == 0 {
			return WhoisResponse{Target: target, Result: "unknown", ASN: nil, ASNName: nil}
		}
		ip = ips[0]
	}

	if isPrivateIP(ip) {
		return WhoisResponse{Target: target, Result: "private IP", ASN: nil, ASNName: nil}
	}

	raw, err := whois.Whois(target)
	var parsed interface{}
	if err != nil {
		parsed = "unknown"
	} else {
		parsed = parseWhois(raw)
	}

	asn, asnName := getASNInfo(ip.String())

	return WhoisResponse{
		Target:  target,
		Result:  parsed,
		ASN:     asn,
		ASNName: asnName,
	}
}

func whoisHandler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("target")
	if target == "" {
		http.Error(w, "Missing 'target' parameter", http.StatusBadRequest)
		return
	}

	result := lookup(target)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func batchHandler(w http.ResponseWriter, r *http.Request) {
	var targets []string
	if err := json.NewDecoder(r.Body).Decode(&targets); err != nil {
		http.Error(w, "Invalid JSON array", http.StatusBadRequest)
		return
	}

	results := make([]WhoisResponse, 0, len(targets))
	for _, target := range targets {
		results = append(results, lookup(target))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func main() {
	http.HandleFunc("/whois", whoisHandler)
	http.HandleFunc("/whois/batch", batchHandler)
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
