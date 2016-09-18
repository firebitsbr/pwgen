// Cryptographically secure password generator.
//
// Copyright (c) 2016 David Peacock - david.j.peacock@gmail.com

package main

import (
	"crypto/rand"
	"flag"
	"fmt"
)

var length = flag.Int("length", 16, "Password length")
var charset = flag.String("charset", "alphanumeric", "Character set: alpha, numeric, alphanumeric")

func main() {
	flag.Parse()

	charSetOption := characterSets(*charset)

	fmt.Println(generator(*length, charSetOption))
}

func characterSets(charSetOption string) string {
	var chars string

	switch {
	case charSetOption == "alpha":
		chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	case charSetOption == "numeric":
		chars = "0123456789"
	case charSetOption == "alphanumeric":
		chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	default:
		chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	}
	return string(chars)
}

func generator(length int, charSet string) string {
	password := make([]byte, length)
	rand.Read(password)

	for k, v := range password {
		password[k] = charSet[v%byte(len(charSet))]
	}
	return string(password)
}
