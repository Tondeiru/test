package main

import (
	"fmt"
	"strings"
)

func between(value string, a string, b string) string {

	posFirst := strings.Index(value, a)
	posLast := strings.Index(value, b)
	posFirstAdjusted := posFirst + len(a)

	if posFirst == -1 {
		return ""
	}
	if posLast == -1 {
		return ""
	}
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

func before(value string, a string) string {
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}

func after(value string, a string) string {
	pos := strings.LastIndex(value, a)
	adjPos := pos + len(a)

	if pos == -1 {
		return ""
	}
	if adjPos >= len(value) {
		return ""
	}
	return value[adjPos:]
}

func main() {
	test := "" //insert smth here

	fmt.Println(between(test, "", "")) //word
	fmt.Println(before(test, ""))      //word
	fmt.Println(after(test, ""))       //word
}
