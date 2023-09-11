package lib

import (
	"fmt"
	"strings"
)

func HasPrefix() {
	fullName := "Fani Alfirdaus"
	prefix := "Fani"

	hasPrefix := strings.HasPrefix(fullName, prefix)

	if hasPrefix {
		fmt.Printf("nama depan saya %s\n", prefix)
	}
}
