package main

import (
	"github.com/notwithering/sgr"
)

const (
	bulkeeError string = sgr.FgHiRed + "error:" + sgr.Reset + " %s\n"
	bulkeeInfo  string = sgr.FgHiBlue + "::" + sgr.Reset + " %s\n"
	bulkeeInput string = ">> "
)
