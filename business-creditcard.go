package creditcard

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Cardtype(ccn string) string {
	ccn = Clean(ccn)
	ccnLen := len(ccn)

	if ccnLen == 0 {
		return ""
	}

	ccType := "Unknown"

	if strings.HasPrefix(ccn, "51") ||
		strings.HasPrefix(ccn, "52") ||
		strings.HasPrefix(ccn, "53") ||
		strings.HasPrefix(ccn, "54") ||
		strings.HasPrefix(ccn, "55") {
		if ccnLen == 16 {
			ccType = "MasterCard"
		}
	} else if strings.HasPrefix(ccn, "4") {
		if ccnLen == 13 || ccnLen == 16 {
			ccType = "Visa"
		}
	} else if strings.HasPrefix(ccn, "34") ||
		strings.HasPrefix(ccn, "37") {
		if ccnLen == 15 {
			ccType = "AmericanExpress"
		}
	}
	return ccType
}

func Validate(ccn string) bool {
	ccn = Clean(ccn)
	ccn = Reverse(ccn)

	even := false
	ccnInt := 0
	for _, r := range ccn {
		c := string(r)
		i, err := strconv.Atoi(c)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}

		if even {
			i *= 2
		}

		if i > 9 {
			ccnInt -= 9
		}

		ccnInt += i

		even = !even
	}

	if ccnInt%10 == 0 {
		return true
	} else {
		return false
	}
}

func Clean(s string) string {
	s = strings.Replace(s, "-", "", -1)
	s = strings.Replace(s, " ", "", -1)
	return s
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
