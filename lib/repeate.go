package lib

import (
	"fmt"
	"strings"
)

func Repeate() {
	nama := "alfi"
	count := 5

	str := strings.Repeat(nama, count)
	fmt.Println(str)
}
