// Copyright 2013 Doug Sparling. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"creditcard"
	"fmt"
)

func main() {
	// true
	fmt.Println(creditcard.Validate("4111111111111111"))
	fmt.Println(creditcard.Validate("4111 1111 1111 1111"))
	fmt.Println(creditcard.Validate("4111-1111-1111-1111"))

	// false
	fmt.Println(creditcard.Validate("1111111111111111"))
	fmt.Println(creditcard.Validate("1111 1111 1111 1111"))
	fmt.Println(creditcard.Validate("1111-1111-1111-1111"))

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
}
