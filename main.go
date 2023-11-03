package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cat input.txt | filter target.txt")
		return
	}

	targetsPath := os.Args[1]

	// Read target domains and IP ranges from file
	targets, err := readTargetsFromFile(targetsPath)
	if err != nil {
		fmt.Println("Error reading target domains:", err)
		return
	}

	// Create a map to store target domains and IP ranges
	targetMap := make(map[string]bool)
	for _, target := range targets {
		targetMap[target] = true
	}

	// Read input domains from standard input (piped in)
	inputDomains, err := readDomainsFromStdin()
	if err != nil {
		fmt.Println("Error reading input domains:", err)
		return
	}

	// Filter input domains
	filteredDomains := []string{}
	for _, domain := range inputDomains {
		if isMatch(domain, targetMap) {
			filteredDomains = append(filteredDomains, domain)
		}
	}

	// Print the filtered domains
	// fmt.Println("Filtered domains:")
	for _, domain := range filteredDomains {
		fmt.Println(domain)
	}
}

// readTargetsFromFile reads target domains and IP ranges from a file and returns them as a slice
func readTargetsFromFile(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	targets := strings.Split(string(content), "\n")
	return targets, nil
}

// readDomainsFromStdin reads domains from standard input (piped in) and returns them as a slice
func readDomainsFromStdin() ([]string, error) {
	var domains []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		domains = append(domains, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return domains, nil
}

// isMatch checks if a domain or IP address matches any of the targets
func isMatch(domainOrIP string, targets map[string]bool) bool {
	for target := range targets {
		if isDomainMatch(domainOrIP, target) || isIPMatch(domainOrIP, target) {
			return true
		}
	}
	return false
}

// isDomainMatch checks if a domain matches the target domain
func isDomainMatch(domain, target string) bool {
	return strings.Contains(domain, target)
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
