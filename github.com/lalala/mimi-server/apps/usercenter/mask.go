package main

import (
	"fmt"
	"strings"
)

func maskPrivacyIdno(idno string) string {
	if len(idno) > 5 {
		return fmt.Sprintf("%s%s%s", idno[:3], strings.Repeat("*", len(idno)-5), idno[len(idno)-2:])
	}
	return idno
}

func maskPrivacyBankcardNo(cardno string) string {
	if len(cardno) < 4 {
		return cardno
	}
	n := len(cardno) - 4
	return fmt.Sprintf("%s%s", strings.Repeat("*", n), cardno[n:])
}

func maskPrivacyRealname(realname string) string {
	ret := ""
	for i, v := range realname {
		if i > 0 {
			ret += "*"
		} else {
			ret += string(v)
		}
	}
	return ret
}
