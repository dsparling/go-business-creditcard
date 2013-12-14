go-business-creditcard
======================

Validate/generate credit card checksums/names.

The go-business-creditcard package is a simple port based on the
[Business::CreditCard](http://search.cpan.org/dist/Business-CreditCard/) Perl module
by [Ivan Kohler](http://search.cpan.org/~ivan/).

## Installation

Simply install the package to your [$GOPATH](http://code.google.com/p/go-wiki/wiki/GOPATH "GOPATH") with the [go tool](http://golang.org/cmd/go/ "go command") from shell:
```bash
$ go get github.com/dsparling/go-business-creditcard
```
Make sure [Git is installed](http://git-scm.com/downloads) on your machine and in your system's `PATH`.

*`go get` installs the latest tagged release*

## Examples

[Example.go](https://github.com/dsparling/go-business-creditcard/blob/master/examples/example.go) is a sort of hello world for go-business-creditcard and should get you started for the barebones necessities of using the package.

	cd examples
	go run example.go

## Validate

	// true
	fmt.Println(creditcard.Validate("4111111111111111"))
	fmt.Println(creditcard.Validate("4111 1111 1111 1111"))
	fmt.Println(creditcard.Validate("4111-1111-1111-1111"))

	// false
	fmt.Println(creditcard.Validate("1111111111111111"))
	fmt.Println(creditcard.Validate("1111 1111 1111 1111"))
	fmt.Println(creditcard.Validate("1111-1111-1111-1111"))

## Cardtype

	// Visa
	fmt.Println(creditcard.Cardtype("4111111111111111"))
	// MasterCard
	fmt.Println(creditcard.Cardtype("5555555555554444"))
	// AmericanExpress
	fmt.Println(creditcard.Cardtype("378282246310005"))
	// DinersClub/Carteblanche
	fmt.Println(creditcard.Cardtype("30569309025904"))
	// Discover
	fmt.Println(creditcard.Cardtype("6011111111111117"))
	// EnRoute
	fmt.Println(creditcard.Cardtype("201400000000009"))
	// JCB
	fmt.Println(creditcard.Cardtype("3530111333300000"))

## GenerateLastDigit

	// Returns '9' - 5276440065421319
	fmt.Println(creditcard.GenerateLastDigit("5276 4400 6542 131"))
