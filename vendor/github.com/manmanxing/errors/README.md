# Forked from https://github.com/pkg/errors

**go1.7+ required**

## Usage

```go
package main

import (
	stderrors "errors"
	"fmt"
	"github.com/manmanxing/errors"
)

func main() {
	err := test3()
	if err != nil {
		fmt.Println(errors.String(err))
		return
	}
}

func test0() error {
	return stderrors.New("original message")
}

func test1() error {
	err := test0()
	if err != nil {
		return errors.Wrap(err,"test1 wrap message")
	}
	return nil
}

func test2() error {
	err := test1()
	if err != nil {
		return errors.Wrap(err, "test2 wrap message")
	}
	return nil
}

func test3() error {
	err := test2()
	if err != nil {
		return errors.Wrap(err,"test3 wrap message")
	}
	return nil
}
```

The result is:
```
$ go run main.go
test3 wrap message: test2 wrap message: test1 wrap message: original message
main.test1
        awesomeProject/main.go:25
main.test2
        awesomeProject/main.go:31
main.test3
        awesomeProject/main.go:39
main.main
        awesomeProject/main.go:11

```