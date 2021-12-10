// Copyright © 2021 Paulo A. P. Júnior
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Based on the "Echo1" program by Alan A. A. Donovan & Brian W. Kernighan.
// https://github.com/adonovan/gopl.io/blob/master/ch1/echo1/main.go

// Ex1-1 prints its name and command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("%s: %s\n", os.Args[0], strings.Join(os.Args[1:], " "))
}
