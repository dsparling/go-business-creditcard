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
}
