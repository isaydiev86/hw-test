package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	msg := stringutil.Reverse("Hello, OTUS!")
	fmt.Println(msg)
}
