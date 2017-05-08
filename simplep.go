package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	helpFlag = flag.Bool("help", false, "Shows the help menu.")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `%s: run several reverse proxies
Usage: %s [flags] {listen|type|target}...
 where type is tcp or http

Flags:
`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}
}

func usageAndExit(reason string) {
	flag.Usage()
	fmt.Fprintf(os.Stderr, "terminating: %s\n", reason)
	os.Exit(1)
}

func errAndExit(err error) {
	usageAndExit(err.Error())
}

func main() {
	flag.Parse()

	if *helpFlag {
		usageAndExit("-help was called")
	}

	if len(flag.Args()) == 0 {
		usageAndExit("no arguments specified")
	}

	for _, target := range flag.Args() {
		splitTarget := strings.Split(target, "|")
		if len(splitTarget) != 3 {
			usageAndExit("target not of type listen|type|target")
		}

		listen := splitTarget[0]
		tunType := splitTarget[1]
		target := splitTarget[2]

		switch tunType {
		case "tcp":
			makeTCPTunnel(listen, target)
		case "http":
			makeHTTPTunnel(listen, target)
		default:
			usageAndExit(fmt.Sprintf("%s isn't a valid type", tunType))
		}
	}

	// now wait forever
	infiniteWait := make(chan bool)
	<-infiniteWait
}
