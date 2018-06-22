package main

import (
	"flag"
	"log"
	"os/exec"
)

// Define the openssl binary used
const openssl = "openssl"

// Shell flags
var (
	domainName               string
	certificateAuthorityName string
)

// Parse flags
func init() {
	flag.StringVar(&domainName, "d", "", "Specifies the domain used for this certificate")
	flag.StringVar(&certificateAuthorityName, "ca", "Self-signed", "Specifies the name of the root certificate authority")
	flag.Parse()
}

func main() {
	if domainName == "" {
		log.Fatal("You need to specify a valid domain name")
	}

	if certificateAuthorityName == "" {
		log.Fatal("You need to specify a valid certificate authority name")
	}

	createCertificateAuthority(certificateAuthorityName, "rootCA.pem", "rootCA.crt")
	createServerCertificate(domainName, "server.pem", "server.crt")
}

func createCertificateAuthority(certificateAuthorityName string, keyFile string, certFile string) {
	err := call(openssl, "genrsa", "-des3", "-out", keyFile, "2048")

	if err != nil {
		log.Fatal("Could not create a key for the certificate authority")
	}

	err = call(openssl, "req", "-x509", "-new", "-nodes", "-key", keyFile, "-sha256", "-days", "3650", "-out", certFile)

	if err != nil {
		log.Fatal("Could not create a certificate for the certificate authority")
	}
}

func createServerCertificate(domain string, keyFile string, certFile string) {
	err := call(openssl, "req", "-new", "-nodes", "-out", "server.csr", "-newkey", "rsa:2048", "-keyout", keyFile)

	if err != nil {
		log.Fatal("Could not create server key")
	}

	// TODO:
	// openssl x509 -req -in server.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out server.crt -days 500 -sha256 -extfile v3.ext
}

func call(command string, arguments ...string) error {
	cmd := exec.Command(command, arguments...)
	return cmd.Run()
}
