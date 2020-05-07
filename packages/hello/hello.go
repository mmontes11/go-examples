package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/mmontes11/go-examples/packages/strings"
)

func main() {
	fmt.Println(strings.Reverse("Hello, world"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
