package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mercari/imageflux-cli/imageflux"
)

func usage() {
	fmt.Println(`Usage:

    imageflux-cli cache.lookup -k $url # lookup cache by key
    imageflux-cli cache.delete -k $url # delete cache by key
`)
	os.Exit(0)
}

func main() {
	if len(os.Args) < 3 {
		usage()
	}
	command := os.Args[1]

	// https://console.imageflux.jp/docs/#api_reference
	switch command {
	case "cache.lookup": // lookup cache by key
		fallthrough
	case "cache.delete": // delete cache by key
		// do nothing
	default:
		usage()
	}

	if err := imageflux.Run(command); err != nil {
		log.Fatal(err)
	}
}
