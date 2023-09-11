package lib

import (
	"fmt"
	"strings"
)

func Join() {
	name := []string{"fani", "alfirdaus"}

	fullName := strings.Join(name, "-")
	fmt.Println(fullName)
}
