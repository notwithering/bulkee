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

	before := getBefore()
	after := getAfter(before)
	remove, create := difference(before, after)
	if !confirm(remove, create) {
		return
	}
	execute(remove, create)
}
