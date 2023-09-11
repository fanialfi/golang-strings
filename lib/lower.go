package lib

import (
	"fmt"
	"strings"
)

func Lower() {
	name := "Fani Alfirdaus"
	fmt.Printf("before\t: \"%s\"\n", name)

	strLower := strings.ToLower(name)
	fmt.Printf("after\t: \"%s\"\n", strLower)
}
