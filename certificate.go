package certificate

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"net"
	"time"
)

// Generate generates new certificates for the given |hosts|.
func Generate(hosts ...string) error {
	validFor := 3650 * 24 * time.Hour
	notBefore := time.Now()
	notAfter := notBefore.Add(validFor)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)

	if err != nil {
		return fmt.Errorf("failed to generate serial number: %s", err)
	}

	// Generate Root CA key
	rootKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		return err
	}

	err = keyToFile("root.key", rootKey)

	if err != nil {
		return err
	}

	// Generate Root CA certificate
	rootTemplate := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
			CommonName:   "Root CA",
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA: true,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &rootTemplate, &rootTemplate, &rootKey.PublicKey, rootKey)

	if err != nil {
		return err
	}

	err = certToFile("root.crt", certBytes)

	if err != nil {
		return err
	}

	// Generate server key
	serverKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		return err
	}

	err = keyToFile("server.key", serverKey)

	if err != nil {
		return err
	}

	// Generate server certificate
	serialNumber, err = rand.Int(rand.Reader, serialNumberLimit)

	if err != nil {
		return err
	}

	serverTemplate := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
			CommonName:   "Localhost Certificate",
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA: false,
	}

	for _, host := range hosts {
		ip := net.ParseIP(host)

		if ip != nil {
			serverTemplate.IPAddresses = append(serverTemplate.IPAddresses, ip)
		} else {
			serverTemplate.DNSNames = append(serverTemplate.DNSNames, host)
		}
	}

	certBytes, err = x509.CreateCertificate(rand.Reader, &serverTemplate, &rootTemplate, &serverKey.PublicKey, rootKey)

	if err != nil {
		return err
	}

	return certToFile("server.crt", certBytes)
}
