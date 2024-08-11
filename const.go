package main

import (
	"github.com/notwithering/sgr"
)

const (
	strError  string = sgr.FgHiRed + "error:" + sgr.Reset + " %s\n"
	strInfo   string = sgr.FgHiBlue + "::" + sgr.Reset + " %s\n"
	strInput  string = ">> "
	strRemove string = " " + sgr.FgHiRed + "-" + sgr.Reset + " %s\n"
	strCreate string = " " + sgr.FgHiGreen + "+" + sgr.Reset + " %s\n"
)
