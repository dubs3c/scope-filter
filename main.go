package main

import (
	"bufio"
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cat input.txt | filter target.txt")
		os.Exit(1)
	}

	targetsPath := os.Args[1]

	// Read target domains and IP ranges from file
	targets, err := readTargetsFromFile(targetsPath)
	if err != nil {
		fmt.Println("Error reading target domains:", err)
		os.Exit(1)
	}

	stdin := os.Stdin

	// Read input domains from standard input (piped in)
	filteredDomains, err := readDomainsFromStdin(targets, stdin)
	if err != nil {
		fmt.Println("Error reading input domains:", err)
		os.Exit(1)
	}

	// Print the filtered domains
	for _, domain := range *filteredDomains {
		fmt.Println(domain)
	}
}

// readTargetsFromFile reads target domains and IP ranges from a file and returns them as a slice
func readTargetsFromFile(filename string) (*map[string]bool, error) {
	targetFileHandle, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer targetFileHandle.Close()

	targetMap := make(map[string]bool)
	scanner := bufio.NewScanner(targetFileHandle)
	for scanner.Scan() {
		targetMap[strings.ToLower((scanner.Text()))] = true
	}

	if err := scanner.Err(); err != nil {
		return &targetMap, err
	}

	return &targetMap, nil
}

// readDomainsFromStdin reads domains from standard input (piped in) and returns them as a slice
func readDomainsFromStdin(targets *map[string]bool, data *os.File) (*[]string, error) {
	var domains []string
	scanner := bufio.NewScanner(data)
	var input string
	for scanner.Scan() {
		input = strings.ToLower(scanner.Text())
		if isMatch(input, targets) {
			domains = append(domains, input)
		}
	}
	if err := scanner.Err(); err != nil {
		return &domains, err
	}
	return &domains, nil
}

// isMatch checks if a domain or IP address matches any of the targets
func isMatch(domainOrIP string, targets *map[string]bool) bool {
	for target := range *targets {
		if isDomainMatch(domainOrIP, target) || isIPMatch(domainOrIP, target) {
			return true
		}
	}
	return false
}

// isDomainMatch checks if a domain matches the target domain
func isDomainMatch(domain, target string) bool {

	// stupid hack
	if !strings.HasPrefix(domain, "https://") && !strings.HasPrefix(domain, "http://") {
		domain = "http://" + domain
	}

	// Parse the domain as a URL
	parsedURL, err := url.Parse(domain)
	if err != nil {
		// Handle parsing error if the input is not a valid URL
		return false
	}

	// Extract the host (domain) from the parsed URL
	host := parsedURL.Host

	// Compare the apex domain with the target
	return strings.Contains(host, target)
}

// isIPMatch checks if an IP address matches the target IP range
func isIPMatch(ip, target string) bool {
	_, targetCIDR, err := net.ParseCIDR(target)
	if err != nil {
		return false
	}

	ipAddr := net.ParseIP(ip)
	if ipAddr == nil {
		return false
	}

	return targetCIDR.Contains(ipAddr)
}
