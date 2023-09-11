# strings

## function `strings.Contains(s, substr string)`

digunakan untuk mengecek apakah string pada parameter kedua merupakan bagian dari string parameter pertama, nilai kembalian dari function ini berupa nilai boolean.

```go
package main

import (
  "fmt"
  "strings"
)

func main(){
  name := "Fani Alfirdaus"
	panggilan := "Alfi"
	isExist := strings.Contains(name, panggilan)

	if isExist {
		fmt.Printf("Hello %s\n", panggilan)
	}
}
```

pada contoh di atas variabel _isExist_ akan berisi `true` karena isi dari variabel _panggilan_ juga merupakan bagian dari variabel _name_.

## function `strings.HasPrefix(s, prefix string)`

digunakan untuk mengecek apakah string pada parameter pertama diawali dengan yang ada pada parameter kedua, nilai kembalian dari function ini berupa boolean.

```go
package main

import (
  "fmt"
  "strings"
)

func main(){
  fullName := "Fani Alfirdaus"
	prefix := "Fani"

	hasPrefix := strings.HasPrefix(fullName, prefix)

	if hasPrefix {
		fmt.Printf("nama depan saya %s\n", prefix)
	}
}
```

## function `strings.HasSuffix(s, suffix string)`

digunakan untuk mengecek apakah isi dari parameter s diakhiri dengan parameter suffix.

```go
package main

import (
  "fmt"
  "strings"
)

func main(){
  fullName := "Fani Alfirdaus"
	suffix := "us"

	isSuffix := strings.HasSuffix(fullName, suffix)

	if isSuffix {
		fmt.Println("OK")
	}
}
```

## function `strings.Count(s, substr string)`

function ini digunakan untuk menghitung seberapa banyak kata atau karaketer pada parameter kedua dari sebuah string pada parameter pertama.

```go
package main

import (
  "fmt"
  "strings"
)

func main(){
  name := "Fani Alfirdaus"
	char := "i"

	howMany := strings.Count(name, char)
	fmt.Printf("huruf \"%s\" ada %d di kata \"%s\"\n", char, howMany, name)
}
```

## function `strings.Index(s, substr string)`

digunakan untuk mencari posisi index parameter kedua dalam string parameter pertama, jika ditemukan lebih dari 1 sub-string maka yang diambil adalah yang pertama.

```go
package main

import (
	"fmt"
	"strings"
)

func main(){
	name := "Fani Alfirdaus"
	char := "i"

	index := strings.Index(name, char)
	fmt.Println(index)
}
```

## function `strings.Replace(s, old, new string, n int)`

digunakan untuk mengganti bagian dari string pada parameter pertama, dimana string pada parameter pertama jika terdapat substring pada parameter kedua, akan diganti dengan substring pada parameter ketiga. 
Jumplah sub-string yang akan di replace bisa ditentukan,apakah hanya 1 sub-string pertama, 2 string, atau seluruhnya dengan men-set nya pada parameter ke empat.

```go
package main

import (
	"fmt"
	"strings"
)

func main(){
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
```

## function `strings.Repeat(s strings, count int)`

digunakan untuk mengulang string pada parameter pertama sebanyak yang telah di tentukan di parameter kedua, akan mengembalikan panic error jika count negatif.

```go
package main

import (
	"fmt"
	"strings"
)

func main(){
	nama := "alfi"
	count := 5

	str := strings.Repeat(nama, count)
	fmt.Println(str)
}
```

## function `strings.Split(s, sep string)`

digunakan untuk memisahkan string pada parameter pertama, dimana separator nya bisa ditentukan sendiri di parameter kedua, dan hasilnya berupa slice string.

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "fani alfirdaus"
	result := strings.Split(name, " ")

	fmt.Println(result)

	for _, elm := range result {
		fmt.Println(elm)
	}
}
```

## function `strings.Join(elms []string, sep string)`

berbeda dengan `strings.Split()` function ini merupakan kebalikannya, digunakan untuk menggabungkan slice string menjadi sebuah string dengan penghubung yang telah ditentukan di parameter kedua.

```go
package main

import (
	"fmt"
	"strings"
)

func main(){
	name := []string{"fani", "alfirdaus"}

	fullName := strings.Join(name, "-")
	fmt.Println(fullName)
}
```

## function `strings.ToLower(s string)`

digunakan untuk mengubah huruf huruf string menjadi kecil.

```go
package main

import (
	"fmt"
	"strings"
)

func main(){
	name := "Fani Alfirdaus"
	fmt.Printf("before\t: \"%s\"\n", name)

	strLower := strings.ToLower(name)
	fmt.Printf("after\t: \"%s\"\n", strLower)
}
```

## function `strings.ToUpper(s string)`

digunakan untuk mengubah huruf huruf string menjadi besar.

```go
package main

import (
	"fmt"
	"strings"
)

func main(){
	name := "Fani Alfirdaus"
	fmt.Printf("before\t: \"%s\"\n", name)

	nameUpper := strings.ToUpper(name)
	fmt.Printf("after\t: \"%s\"\n", nameUpper)
}
```