go-business-creditcard
======================

The go-business-creditcard package is a simple port of the
[Business::CreditCard](http://search.cpan.org/dist/Business-CreditCard/) Perl module
by [Ivan Kohler](http://search.cpan.org/~ivan/).

Note - this code is under development and not yet complete.

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

[Example.go][example] is a sort of hello world for go-business-creditcard and should get you started for the barebones necessities of using the package.

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
