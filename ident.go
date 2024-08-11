package main

import (
	"strings"
)

func getIdentifier(name string, isDir bool) string {
	var ident string

	if strings.ContainsAny(name, " \t\n\\\"'~!@#$%^&*()<>?") {
		ident += "'"
		ident += name
		ident += "'"
	} else {
		ident = name
	}

	if isDir {
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

func cleanIdentifier(id string) string {
	return getIdentifier(trimIdentifier(id))
}
