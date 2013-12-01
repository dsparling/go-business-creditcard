go-business-creditcard
======================

Validate/generate credit card checksums/names.

The go-business-creditcard package is a simple port based on the
[Business::CreditCard](http://search.cpan.org/dist/Business-CreditCard/) Perl module
by [Ivan Kohler](http://search.cpan.org/~ivan/).

## Installation

go-business-creditcard is built using the Go tool. The tool assumes the code will be in a folder $GOPATH/src/go-business-creditcard.

	cd $GOPATH/src
	git clone git://github.com/dsparling/go-business-creditcard creditcard
	cd creditcard
	go install

Confirm the install has created redis.a in your $GOPATH/pkg/<arch> folder:

	ls -l $GOPATH/pkg/"$GOOS"_"$GOARCH"/creditcard.a

e.g. on my Mac OS X (64b)

	ls -l <my-gopath>/pkg/darwin_amd64

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
