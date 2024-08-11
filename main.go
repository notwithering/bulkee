package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Usage = func() {
		fmt.Println("usage: bulkee <dir>")
	}
	flag.Parse()

	dir := flag.Arg(0)
	if dir == "" {
		dir = "."
	}

	before := getBefore(dir)
	after := getAfter(before)
	remove, create := difference(before, after)
	if !confirm(remove, create) {
		return
	}
	execute(dir, remove, create)
}
