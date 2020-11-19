# Read File

readfile is a simple package to help reading files

## How use 

```bash
go get github.com/caioreix/readfile
```

### File
Read an entire file and populate a string

```go
package main

import (
	"fmt"
	"log"

	"github.com/caioreix/readfile"
)

func main() {
	content, err := readfile.File("path/to/my/file")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(content)
}
```

### ByLine
Read a file line by line

```go
package main

import (
	"fmt"
	"log"

	"github.com/caioreix/readfile"
)

func main() {
	lines, err := readfile.ByLine("path/to/my/file")
	if err != nil {
		log.Fatal(err)
	}

	for i, l := range lines {
		fmt.Printf("%d. %s\n", i, l)
	}
}
```

### ByWord
Read a file word by word

```go
package main

import (
	"fmt"
	"log"

	"github.com/caioreix/readfile"
)

func main() {
	words, err := readfile.ByWord("path/to/my/file")
	if err != nil {
		log.Fatal(err)
	}

	for i, w := range words {
		fmt.Printf("%d. %s\n", i, l)
	}
}
```