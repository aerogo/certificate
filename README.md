# certificate

[![Godoc][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][coverage-image]][coverage-url]
[![Patreon][patreon-image]][patreon-url]

Creates self-signed certificates for local development using ECDSA and x509 Version 3. This will give you a green lock in modern versions of Chrome.

## Installation

```shell
go get -u github.com/aerogo/certificate/...
```

## Usage

### Single domain

```shell
makecert -host local.example.com
```

### Multi domain

```shell
makecert -host local.example.com,localhost,127.0.0.1
```

## Output

The tool will create 4 files in the current directory:

* **root.crt** (import this as a trusted authority into your browser)
* **root.key**
* **server.crt** (certificate used in your server)
* **server.key** (private key used in your server)

## Style

Please take a look at the [style guidelines](https://github.com/akyoto/quality/blob/master/STYLE.md) if you'd like to make a pull request.

## Sponsors

| [![Scott Rayapoullé](https://avatars3.githubusercontent.com/u/11772084?s=70&v=4)](https://github.com/soulcramer) | [![Eduard Urbach](https://avatars2.githubusercontent.com/u/438936?s=70&v=4)](https://twitter.com/eduardurbach) |
| --- | --- |
| [Scott Rayapoullé](https://github.com/soulcramer) | [Eduard Urbach](https://eduardurbach.com) |

Want to see [your own name here?](https://www.patreon.com/eduardurbach)

[godoc-image]: https://godoc.org/github.com/aerogo/certificate?status.svg
[godoc-url]: https://godoc.org/github.com/aerogo/certificate
[report-image]: https://goreportcard.com/badge/github.com/aerogo/certificate
[report-url]: https://goreportcard.com/report/github.com/aerogo/certificate
[tests-image]: https://cloud.drone.io/api/badges/aerogo/certificate/status.svg
[tests-url]: https://cloud.drone.io/aerogo/certificate
[coverage-image]: https://codecov.io/gh/aerogo/certificate/graph/badge.svg
[coverage-url]: https://codecov.io/gh/aerogo/certificate
[patreon-image]: https://img.shields.io/badge/patreon-donate-green.svg
[patreon-url]: https://www.patreon.com/eduardurbach
