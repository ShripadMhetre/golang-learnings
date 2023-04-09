package main

import (
	"os"

	"github.com/shripadmhetre/golang-learnings/di"
	"github.com/shripadmhetre/golang-learnings/oops"
)

func main() {
	arg1 := os.Args[1]

	// use arg1 to decide which packages to call
	if arg1 == "di" {
		di.Main()
	} else if arg1 == "oops" {
		oops.Main()
	}
}
