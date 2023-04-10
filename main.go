package main

import (
	"github.com/shripadmhetre/golang-learnings/beforedi"
	"github.com/shripadmhetre/golang-learnings/oops"
	"github.com/shripadmhetre/golang-learnings/wiredemo"
	"os"

	"github.com/shripadmhetre/golang-learnings/di"
)

func main() {
	arg1 := os.Args[1]

	if arg1 == "oops" {
		oops.Main()
	} else if arg1 == "beforedi" {
		beforedi.Main()
	} else if arg1 == "di" {
		di.Main()
	} else if arg1 == "wire" {
		wiredemo.Main()
	}
}
