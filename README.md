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