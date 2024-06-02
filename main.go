package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"slices"
	"strings"
	"time"

	"github.com/gobuffalo/envy"
)

func main() {
	flag.Parse()

	var workingDirectory string = flag.Arg(0)
	if workingDirectory == "" {
		workingDirectory = "."
	}

	entries, err := os.ReadDir(workingDirectory)
	if err != nil {
		fmt.Println(err)
		return
	}

	var before []string

	for _, entry := range entries {
		before = append(before, getIdentifier(entry))
	}

	tmpFile, err := os.CreateTemp(os.TempDir(), "bulkee")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(tmpFile.Name())

	tmpFile.Write([]byte(strings.Join(before, "\n") + "\n"))
	tmpFile.Close()

	editor := envy.Get("EDITOR", "vim")

	cmd := exec.Command(editor, tmpFile.Name())

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()

	tmpFile, err = os.Open(tmpFile.Name())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tmpFile.Close()

	var after []string

	scanner := bufio.NewScanner(tmpFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		after = append(after, line)
	}

	var remove []string
	for _, id := range before {
		if !slices.Contains(after, id) {
			remove = append(remove, id)
		}
	}

	var create []string
	for _, id := range after {
		if !slices.Contains(before, id) {
			create = append(create, id)
		}
	}

	fmt.Println("Remove:")
	fmt.Println(strings.Join(remove, ", "))

	fmt.Println("Create:")
	fmt.Println(strings.Join(create, ", "))

	fmt.Println("Executing in 3 seconds...")
	time.Sleep(3 * time.Second)

	for _, id := range remove {
		name, _ := trimIdentifier(id)
		if err := os.RemoveAll(name); err != nil {
			fmt.Println(err)
		}
	}

	for _, id := range create {
		name, isDir := trimIdentifier(id)
		if isDir {
			if err := os.Mkdir(name, os.ModePerm); err != nil {
				fmt.Println(err)
			}
		} else {
			if _, err := os.Create(name); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func getIdentifier(entry fs.DirEntry) string {
	var ident string
	name := entry.Name()

	if strings.ContainsAny(name, " \t\n\\\"'~!@#$%^&*()<>?") {
		ident += "'"
		ident += entry.Name()
		ident += "'"
	} else {
		ident = entry.Name()
	}

	if entry.IsDir() {
		ident += "/"
	}

	return ident
}

func trimIdentifier(id string) (name string, isDir bool) {
	name, isDir = strings.CutSuffix(id, "/")
	name = strings.TrimPrefix(name, "'")
	name = strings.TrimSuffix(name, "'")
	return name, isDir
}
