package main

import (
	"github.com/bean-du/emailFilter/read"
)

func main() {
	f := read.GetNewFiles(200)
	f.Read()
	f.Filter()
	f.Write()
}

