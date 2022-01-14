package main

import (
	"fmt"
	"imagine-flow/vars"
	"imagine-flow/web"
)

var Version string

func main() {
	fmt.Println("Imagine Flow Version", vars.Version)
	web.StartWebs()
}
