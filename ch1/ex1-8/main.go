// Copyright © 2021 Paulo A. P. Júnior
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Based on the "Fetch" program by Alan A. A. Donovan & Brian W. Kernighan.
// https://github.com/adonovan/gopl.io/blob/master/ch1/fetch/main.go

// Ex1-8 prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fecth: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		defer resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fecth: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
