package lib

import (
	"fmt"
	"strings"
)

func Replace() {
	text := "banana"
	find := "a"
	replaceWith := "o"

	newText1 := strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1)

	newText2 := strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2)

	newText3 := strings.Replace(text, find, replaceWith, 3)
	fmt.Println(newText3)

	newText4 := strings.Replace(text, find, replaceWith, 4)
	fmt.Println(newText4)
}
