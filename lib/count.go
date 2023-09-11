package lib

import (
	"fmt"
	"strings"
)

func Count() {
	name := "Fani Alfirdaus"
	char := "i"

	howMany := strings.Count(name, char)
	fmt.Printf("huruf \"%s\" ada %d di kata \"%s\"\n", char, howMany, name)
}
