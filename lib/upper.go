package lib

import (
	"fmt"
	"strings"
)

func Upper() {
	name := "Fani Alfirdaus"
	fmt.Printf("before\t: \"%s\"\n", name)

	nameUpper := strings.ToUpper(name)
	fmt.Printf("after\t: \"%s\"\n", nameUpper)
}
