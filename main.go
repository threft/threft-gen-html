package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/jessevdk/go-flags"
	"github.com/threft/threft/tidm"
	"os"
)

var options struct {
	Debugging bool `short:"d" long:"debug" description:"Enable logging of debug messages to StdOut"`
}

func main() {
	args, err := flags.Parse(&options)
	if err != nil {
		flagError := err.(*flags.Error)
		if flagError.Type == flags.ErrHelp {
			fmt.Println("Example: threft-gen-go -p json -p binary")
			fmt.Println()
			return
		}
		if flagError.Type == flags.ErrUnknownFlag {
			fmt.Println("Use --help to view all available options.")
			return
		}
		fmt.Printf("Error parsing flags: %s\n", err)
		return
	}
	if len(args) > 0 {
		fmt.Printf("Unknown argument '%s'.", args[0])
		return
	}

	if options.Debugging {
		fmt.Println("threft-gen-html started, reading TIDM from stdin.")
	}

	t, err := tidm.ReadFrom(os.Stdin)
	if err != nil {
		fmt.Printf("Error reading TIDM from stdin: %s\n", err)
		return
	}
	spew.Dump(t)
}
