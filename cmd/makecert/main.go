package main

import (
	"flag"
	"os"
	"strings"

	"github.com/aerogo/certificate"
)

var host string

func init() {
	flag.StringVar(&host, "host", "", "Comma-separated hostnames and IPs to generate a certificate for")
	flag.Parse()
}

func main() {
	if host == "" {
		flag.Usage()
		os.Exit(1)
	}

	err := certificate.Generate(strings.Split(host, ",")...)

	if err != nil {
		panic(err)
	}
}
