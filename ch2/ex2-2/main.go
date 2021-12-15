package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		parseAndShow(arg)
	}

	// Does not have args. Use standard input
	if len(os.Args) == 1 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			parseAndShow(input.Text())
		}
	}
}

func parseAndShow(input string) {
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gconv: %v\n", err)
		return
	}
	c := Celsius(num)
	f := Fahrenheit(num)
	ft := Ft(num)
	km := Km(num)
	lb := Lb(num)
	kg := Kg(num)

	fmt.Printf("%s = %s, %s = %s\n", c, CToF(c), f, FToC(f))
	fmt.Printf("%s = %s, %s = %s\n", ft, FtToKm(ft), km, KmToFt(km))
	fmt.Printf("%s = %s, %s = %s\n\n", lb, LbToKg(lb), kg, KgToLb(kg))
}
