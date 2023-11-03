package main

import (
	"os"
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {

	// Read target domains and IP ranges from file
	targets, err := readTargetsFromFile("test/target.txt")
	if err != nil {
		t.Fatal("Error reading target domains:", err)
	}

	inputHandle, err := os.Open("test/input.txt")
	if err != nil {
		t.Fatal("Error reading input file", err)
	}
	defer inputHandle.Close()

	// Read input domains from standard input (piped in)
	filteredDomains, err := readDomainsFromStdin(targets, inputHandle)
	if err != nil {
		t.Fatal("Error reading input domains:", err)
	}

	if len(*filteredDomains) == 0 {
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
	if !reflect.DeepEqual(expectedOutput, *filteredDomains) {
		t.Errorf("expected %v, got %v", expectedOutput, filteredDomains)
	}
}
