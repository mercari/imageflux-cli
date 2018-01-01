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

    imageflux-cli signature -s $secret -p $path # calculate signature of signed url
`)
	os.Exit(0)
}

func main() {
	if len(os.Args) < 3 {
		usage()
	}
	command := os.Args[1]

	// https://console.imageflux.jp/docs/
	switch command {
	case "cache.lookup": // lookup cache by key
		fallthrough
	case "cache.delete": // delete cache by key
		if err := imageflux.Issue(command); err != nil {
			log.Fatal(err)
		}
	case "signature": // calculate signature of signed url
		signature, err := imageflux.Signature(command)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(signature)
	default:
		usage()
	}

}
