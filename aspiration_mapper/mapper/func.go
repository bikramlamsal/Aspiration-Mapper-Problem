package mapper

import (
	"fmt"
	"strings"
	"unicode"
)

// CapitalizeEveryThirdAlphanumericChar that capitalizes *only* every third alphanumeric character of a string, so that if I hand you
func CapitalizeEveryThirdAlphanumericChar(s string) string {
	s = strings.ToLower(s)
	fmt.Println(s)
	idx := 1
	var resp []rune
	for _, v := range s {
		var x rune
		if unicode.IsLetter(v) {
			if idx%3 == 0 {
				x = unicode.ToUpper(v)
			} else {
				x = v
			}
			idx++
		} else {
			x = v
		}
		fmt.Println(string(x))
		resp = append(resp, x)
	}

	return string(resp)
}
