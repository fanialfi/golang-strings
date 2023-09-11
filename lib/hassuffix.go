package lib

import (
	"fmt"
	"strings"
)

func HasSuffix() {
	fullName := "Fani Alfirdaus"
	suffix := "us"

	isSuffix := strings.HasSuffix(fullName, suffix)

	if isSuffix {
		fmt.Println("OK")
	}
}
