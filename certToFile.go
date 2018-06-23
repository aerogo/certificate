package certificate

import (
	"encoding/pem"
	"os"
)

// certToFile writes a PEM serialization of |certBytes| to a new file called |fileName|.
func certToFile(fileName string, certBytes []byte) error {
	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	return pem.Encode(file, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})
}
