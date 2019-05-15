package certificate_test

import (
	"os"
	"testing"

	"github.com/aerogo/certificate"
)

func TestGenerate(t *testing.T) {
	defer os.Remove("root.key")
	defer os.Remove("root.crt")
	defer os.Remove("server.key")
	defer os.Remove("server.crt")

	err := certificate.Generate("example.com", "localhost", "127.0.0.1")

	if err != nil {
		panic(err)
	}
}
