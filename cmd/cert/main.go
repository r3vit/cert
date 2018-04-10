package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/genkiroid/cert"
)

var version = ""

func main() {
	var format string
	var template string
	var skipVerify bool
	var utc bool
	var showVersion bool

	flag.StringVar(&format, "f", "simple table", "Output format. md: as markdown, json: as JSON. ")
	flag.StringVar(&format, "format", "simple table", "Output format. md: as markdown, json: as JSON. ")
	flag.StringVar(&template, "t", "", "Output format as Go template string or Go template file path.")
	flag.StringVar(&template, "template", "", "Output format as Go template string or Go template file path.")
	flag.BoolVar(&skipVerify, "k", false, "Skip verification of server's certificate chain and host name.")
	flag.BoolVar(&skipVerify, "skip-verify", false, "Skip verification of server's certificate chain and host name.")
	flag.BoolVar(&utc, "u", false, "Set UTC to timezone.")
	flag.BoolVar(&utc, "utc", false, "Set UTC to timezone.")
	flag.BoolVar(&showVersion, "v", false, "Show version.")
	flag.BoolVar(&showVersion, "version", false, "Show version.")
	flag.Parse()

	if showVersion {
		fmt.Println("cert version ", version)
		return
	}

	var certs cert.Certs
	var err error

	cert.SkipVerify = skipVerify
	cert.UTC = utc

	certs, err = cert.NewCerts(flag.Args())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if template == "" {
		switch format {
		case "md":
			fmt.Printf("%s", certs.Markdown())
		case "json":
			fmt.Printf("%s", certs.JSON())
		default:
			fmt.Printf("%s", certs)
		}
		return
	}

	if err := cert.SetUserTempl(template); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s", certs)
}
