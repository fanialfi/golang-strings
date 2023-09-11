package lib

import (
	"fmt"
	"strings"
)

func Contains() {
	name := "Fani Alfirdaus"
	panggilan := "Alfi"
	isExist := strings.Contains(name, panggilan)

	if isExist {
		fmt.Printf("Hello %s\n", panggilan)
	}
}
