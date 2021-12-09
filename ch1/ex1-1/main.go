package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("%s: %s\n", os.Args[0], strings.Join(os.Args[1:], " "))
}
