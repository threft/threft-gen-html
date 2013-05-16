package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/threft/threft/tidm"
	"os"
)

func main() {
	fmt.Println("Hello. This is threft-gen-html")
	t, err := tidm.ReadFrom(os.Stdin)
	if err != nil {
		fmt.Printf("Error reading TIDM from stdin: %s\n", err)
		return
	}
	spew.Dump(t)
}
