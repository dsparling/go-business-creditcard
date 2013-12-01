package creditcard

import (
        "fmt"
        "os"
        "strconv"
)

func Validate(ccn string) bool {
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

    if ccnInt % 10 == 0 {
        return true
    } else {
        return false
    }
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
