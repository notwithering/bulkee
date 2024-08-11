package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func confirm(remove, create []string) bool {
	if len(remove)+len(create) == 0 {
		return false
	}

	if len(remove) > 0 {
		fmt.Printf(strInfo, "Entries will be removed:")
		for _, id := range remove {
			fmt.Printf(strRemove, id)
		}

		if len(create) > 0 {
			fmt.Print("\n")
		}
	}

	if len(create) > 0 {
		fmt.Printf(strInfo, "Entries will be created:")
		for _, id := range create {
			fmt.Printf(strCreate, id)
		}
	}

	fmt.Print("\n")

	fmt.Printf(strInfo, "Proceed with execution? [Y/n]")
	fmt.Print(strInput)

	yes, err := yesno(true)
	if err != nil {
		fmt.Printf(strError, err)
		os.Exit(1)
	}

	return yes
}

func execute(dir string, remove, create []string) {
	for _, id := range remove {
		name, _ := trimIdentifier(id)
		if err := os.RemoveAll(filepath.Join(dir, name)); err != nil {
			fmt.Println(err)
		}
	}

	for _, id := range create {
		name, isDir := trimIdentifier(id)
		if isDir {
			if err := os.Mkdir(filepath.Join(dir, name), os.ModePerm); err != nil {
				fmt.Println(err)
			}
		} else {
			if _, err := os.Create(filepath.Join(dir, name)); err != nil {
				fmt.Println(err)
			}
		}
	}
}
