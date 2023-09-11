package lib

import (
	"fmt"
	"strings"
)

func Index() {
	name := "Fani Alfirdaus"
	char := "i"

	index := strings.Index(name, char)
	fmt.Println(index)
}
