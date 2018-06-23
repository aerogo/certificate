# certificate

Creates self-signed certificates for local development using ECDSA and x509 Version 3. This will give you a green lock in modern versions of Chrome.

## Installation

```
go get github.com/aerogo/certificate/cmd/makecert
```

## Usage

### Single domain

```
makecert -host local.example.com
```

### Multi domain

```
makecert -host local.example.com,localhost,127.0.0.1
```

## Output

The tool will create 4 files in the current directory:

* **root.crt** (import this as a trusted authority into your browser)
* **root.key**
* **server.crt** (certificate used in your server)
* **server.key** (private key used in your server)