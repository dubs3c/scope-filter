package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {

	// Read target domains and IP ranges from file
	targets, err := readTargetsFromFile("test/target.txt")
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
	// In normal use case, this would be read from stdin
	inputDomains, err := readTargetsFromFile("test/input.txt")
	if err != nil {
		fmt.Println("Error reading input domains:", err)
		t.Fatal(err)
	}

	if len(inputDomains) == 0 {
		t.Fatal("inputDomains is empty")
	}

	// Filter input domains
	filteredDomains := []string{}
	for _, domain := range inputDomains {
		if isMatch(domain, targetMap) {
			filteredDomains = append(filteredDomains, domain)
		}
	}

	if len(filteredDomains) == 0 {
		t.Fatal("filteredDomains is empty")
	}

	expectedOutput := []string{
		"bookie.dubell.io",
		"shit.infd.pw",
		"hidden.c.collab.dubell.io",
		"https://test.dubell.io/this/is/a/path",
		"https://shit.infd.pw/admin.panel.html",
		"192.168.1.34",
		"192.168.1.156",
	}

	// Compare the program's output with the expected output
	if !reflect.DeepEqual(expectedOutput, filteredDomains) {
		t.Errorf("expected %v, got %v", expectedOutput, filteredDomains)
	}
}
