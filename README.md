# goshell: Go Shell 
A simple wrapper over os/exec to run shell command using `"sh -c"`.

## Installation

```bash
go install "github.com/miteshbsjat/goshell"
```

## Sample Run

### Simple shell command run

* Go Code

```go
package main

import (
	"fmt"

	"github.com/miteshbsjat/goshell"
)

func main() {
	output, _ := goshell.RunCommand("date")
	fmt.Println(output)
}
```

* Sample output

```bash
Thu Mar 23 07:59:15 PM IST 2023
```


### Iterate output

* Go code

```go
package main

import (
	"fmt"

	"github.com/miteshbsjat/goshell"
)

func main() {
	output, _ := goshell.RunCommand("for i in `seq 1 10`; do echo $i; done")
	for _, line := range goshell.SplitLines(output) {
		fmt.Printf("%s\n", line)
	}
}
```

* Output

```bash
1
2
3
4
5
6
7
8
9
10
```

* More examples are in [mitesh_test](mitesh_test.go)