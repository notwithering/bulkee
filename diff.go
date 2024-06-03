package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
)

var workingDirectory string

func getBefore() []string {
	workingDirectory = flag.Arg(0)
	if workingDirectory == "" {
		workingDirectory = "."
	}

	entries, err := os.ReadDir(workingDirectory)
	if err != nil {
		fmt.Printf(bulkeeError, err)
		os.Exit(1)
	}

	var before []string

	for _, entry := range entries {
		before = append(before, getIdentifier(entry.Name(), entry.IsDir()))
	}

	return before
}

func getAfter(before []string) []string {
	name := edit(before)

	tmpFile, err := os.Open(name)
	if err != nil {
		fmt.Printf(bulkeeError, err)
		os.Exit(1)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	var after []string

	scanner := bufio.NewScanner(tmpFile)
	for scanner.Scan() {
		id := scanner.Text()

		if id == "" {
			continue
		}

		after = append(after, cleanIdentifier(id))
	}

	return after
}

func difference(before, after []string) (remove, create []string) {
	for _, id := range before {
		if !slices.Contains(after, id) {
			remove = append(remove, id)
		}
	}

	for _, id := range after {
		if !slices.Contains(before, id) {
			create = append(create, id)
		}
	}

	return remove, create
}
