package main

import (
	"flag"
)

func main() {
	flag.Parse()

	before := getBefore()
	after := getAfter(before)
	remove, create := difference(before, after)
	if !confirm(remove, create) {
		return
	}
	execute(remove, create)
}
