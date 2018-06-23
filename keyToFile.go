package certificate

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// keyToFile writes a PEM serialization of |key| to a new file called |fileName|.
func keyToFile(fileName string, key *ecdsa.PrivateKey) error {
	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	b, err := x509.MarshalECPrivateKey(key)

	if err != nil {
		return err
	}

	return pem.Encode(file, &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: b,
	})
}
