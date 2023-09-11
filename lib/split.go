package lib

import (
	"fmt"
	"strings"
)

func Split() {
	name := "fani alfirdaus"
	result := strings.Split(name, " ")

	fmt.Println(result)

	for _, elm := range result {
		fmt.Println(elm)
	}
}
