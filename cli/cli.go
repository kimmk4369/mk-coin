package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/kimmk4369/mk-coin/explorer"
	"github.com/kimmk4369/mk-coin/rest"
)

func usage() {
	fmt.Printf("Welcome to MK코인\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port:		set the PORT of the server\n")
	fmt.Printf("-mode:		Choose between 'html' and 'rest'\n\n")
	runtime.Goexit()
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		// start rest api
		rest.Start(*port)
	case "html":
		// start html explorer
		explorer.Start(*port)
	default:
		usage()
	}
}
