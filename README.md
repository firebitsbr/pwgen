# pwgen

[![Build Status](https://travis-ci.org/davidjpeacock/pwgen.svg?branch=master)](https://travis-ci.org/davidjpeacock/pwgen)
[![Go Report Card](https://goreportcard.com/badge/github.com/davidjpeacock/pwgen)](https://goreportcard.com/report/github.com/davidjpeacock/pwgen)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/davidjpeacock/pwgen/master/LICENSE)

Cryptographically secure password generator implemented in Golang.

Supports arbitrary password length, and alpha, numeric, and alphanumeric character sets.

## Installation

```
go get github.com/davidjpeacock/pwgen
```

## Usage

```
Usage of pwgen:
  -charset string
        Character set: alpha, numeric, alphanumeric (default "alphanumeric")
  -length int
        Password length (default 16)
  -v    Prints pwgen version
```

## Example

```
$ pwgen -length 20 -charset alpha
FPxvdWzfhPAmsCnKSGAd
```

## License

 MIT
  
## Contributing

  Pull requests welcome.
