// Cryptographically secure password generator.
//
// Copyright (c) 2016 David Peacock - david.j.peacock@gmail.com

package main

import (
	"crypto/rand"
	"flag"
	"fmt"
)

const version = "0.0.3"

var length = flag.Int("length", 16, "Password length")
var charset = flag.String("charset", "alphanumeric", "Character set: alpha, numeric, alphanumeric, special, qwerty")
var v = flag.Bool("v", false, "Prints pwgen version")

func main() {
	flag.Parse()

	if *v {
		fmt.Println("pwgen version: " + version)
		return
	}

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
	case charSetOption == "special":
		chars = "~`!@#$%^&*()-=_+{}[];':\"|<>,./?\\ "
	case charSetOption == "qwerty":
		chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789~`!@#$%^&*()-=_+{}[];':\"|<>,./?\\ "
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
